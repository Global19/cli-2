package expander

import (
	"regexp"
	"strings"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/constraints"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/print"
	secretsapi "github.com/ActiveState/cli/internal/secrets-api"
	"github.com/ActiveState/cli/pkg/projectfile"
)

var (
	// FailExpandVariable identifies a failure during variable expansion.
	FailExpandVariable = failures.Type("expander.fail.expandvariable", failures.FailUser)

	// FailExpandVariableBadCategory identifies a variable expansion failure due to a bad variable category.
	FailExpandVariableBadCategory = failures.Type("expander.fail.expandvariable.badcategory", FailExpandVariable)

	// FailExpandVariableBadName identifies a variable expansion failure due to a bad variable name.
	FailExpandVariableBadName = failures.Type("expander.fail.expandvariable.badName", FailExpandVariable)

	// FailExpandVariableRecursion identifies a variable expansion failure due to infinite recursion.
	FailExpandVariableRecursion = failures.Type("expander.fail.expandvariable.recursion", FailExpandVariable)

	// FailExpanderBadName is used when an Expanders name is invalid.
	FailExpanderBadName = failures.Type("expander.fail.expander.badName", failures.FailVerify)

	// FailExpanderNoFunc is used when no handler func is found for an Expander.
	FailExpanderNoFunc = failures.Type("expander.fail.expander.noFunc", failures.FailVerify)

	// FailVarNotFound is used when no handler func is found for an Expander.
	FailVarNotFound = failures.Type("expander.fail.vars.notfound", FailExpandVariable)
)

var lastFailure *failures.Failure

// Failure retrieves the latest failure
func Failure() *failures.Failure {
	return lastFailure
}

// Expand will detect the active project and invoke ExpandFromProject with the given string
func Expand(s string) string {
	return ExpandFromProject(s, projectfile.Get())
}

// ExpandFromProject searches for $category.name-style variables in the given
// string and substitutes them with their contents, derived from the given
// project, and subject to the given constraints (if any).
func ExpandFromProject(s string, p *projectfile.Project) string {
	return limitExpandFromProject(0, s, p)
}

// limitExpandFromProject limits the depth of an expansion to avoid infinite expansion of a value.
func limitExpandFromProject(depth int, s string, p *projectfile.Project) string {
	lastFailure = nil
	if depth > constants.ExpanderMaxDepth {
		lastFailure = FailExpandVariableRecursion.New("error_expand_variable_infinite_recursion", s)
		print.Warning(lastFailure.Error())
		return ""
	}

	regex := regexp.MustCompile("\\${?\\w+\\.[\\w-]+}?")
	expanded := regex.ReplaceAllStringFunc(s, func(variable string) string {
		components := strings.Split(strings.Trim(variable, "${}"), ".")
		category, name := components[0], components[1]
		var value string

		if expanderFn, foundExpander := expanderRegistry[category]; foundExpander {
			var failure *failures.Failure

			if value, failure = expanderFn(name, p); failure != nil {
				lastFailure = FailExpandVariableBadName.New("error_expand_variable_project_unknown_name", variable, failure.Error())
				print.Warning(lastFailure.Error())
			}
		} else {
			lastFailure = FailExpandVariableBadCategory.New("error_expand_variable_project_unknown_category", variable, category)
			print.Warning(lastFailure.Error())
		}

		if value != "" {
			value = limitExpandFromProject(depth+1, value, p)
		}
		return value
	})

	return expanded
}

// Func defines an Expander function which can expand the name for a category. An Expander expects the name
// to be expanded along with the project-file definition. It will return the expanded value of the name
// or a Failure if expansion was unsuccessful.
type Func func(name string, project *projectfile.Project) (string, *failures.Failure)

// PlatformExpander expends metadata about the current platform.
func PlatformExpander(name string, project *projectfile.Project) (string, *failures.Failure) {
	for _, platform := range project.Platforms {
		if !constraints.PlatformMatches(platform) {
			continue
		}

		switch name {
		case "name":
			return platform.Name, nil
		case "os":
			return platform.Os, nil
		case "version":
			return platform.Version, nil
		case "architecture":
			return platform.Architecture, nil
		case "libc":
			return platform.Libc, nil
		case "compiler":
			return platform.Compiler, nil
		default:
			return "", FailExpandVariableBadName.New("error_expand_variable_project_unrecognized_platform_var", name)
		}
	}
	return "", nil
}

// EventExpander expands events defined in the project-file.
func EventExpander(name string, project *projectfile.Project) (string, *failures.Failure) {
	var value string
	for _, event := range project.Events {
		if event.Name == name && !constraints.IsConstrained(event.Constraints) {
			value = event.Value
			break
		}
	}
	return value, nil
}

// ScriptExpander expands scripts defined in the project-file.
func ScriptExpander(name string, project *projectfile.Project) (string, *failures.Failure) {
	var value string
	for _, script := range project.Scripts {
		if script.Name == name && !constraints.IsConstrained(script.Constraints) {
			value = script.Value
			break
		}
	}
	return value, nil
}

// VarExpander takes car of expanding user defined variables
type VarExpander struct {
	secretsClient   *secretsapi.Client
	secretsExpander SecretFunc
}

// Expand is the main expander function
func (e *VarExpander) Expand(name string, projectFile *projectfile.Project) (string, *failures.Failure) {
	var variable *projectfile.Variable
	for _, varcheck := range projectFile.Variables {
		if varcheck.Name == name && !constraints.IsConstrained(varcheck.Constraints) {
			variable = varcheck
			break
		}
	}

	if variable == nil {
		return "", FailVarNotFound.New("variables_expand_err_undefined", name)
	}

	if variable.Value.StaticValue != nil {
		return *variable.Value.StaticValue, nil
	}

	return e.secretsExpander(variable, projectFile)
}

// NewVarExpander creates an Expander which can retrieve and decrypt stored user secrets.
func NewVarExpander(secretsClient *secretsapi.Client) Func {
	secretsExpander := NewSecretExpander(secretsClient)
	expander := &VarExpander{secretsClient, secretsExpander.Expand}
	return expander.Expand
}

// NewVarPromptingExpander creates an Expander which can retrieve and decrypt stored user secrets. Additionally,
// it will prompt the user to provide a value for a secret -- in the event none is found -- and save the new
// value with the secrets service.
func NewVarPromptingExpander(secretsClient *secretsapi.Client) Func {
	secretsExpander := NewSecretExpander(secretsClient)
	expander := &VarExpander{secretsClient, secretsExpander.ExpandWithPrompt}
	return expander.Expand
}