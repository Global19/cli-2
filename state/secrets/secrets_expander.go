package secrets

import (
	"strings"

	"github.com/ActiveState/cli/internal/api/models"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/keypairs"
	"github.com/ActiveState/cli/internal/organizations"
	"github.com/ActiveState/cli/internal/projects"
	secretsapi "github.com/ActiveState/cli/internal/secrets-api"
	secretsModels "github.com/ActiveState/cli/internal/secrets-api/models"
	"github.com/ActiveState/cli/internal/variables"
	"github.com/ActiveState/cli/pkg/projectfile"
)

// NewExpander creates an ExpanderFunc which can decrypt stored user secrets.
func NewExpander(secretsClient *secretsapi.Client) variables.ExpanderFunc {
	return func(name string, projectFile *projectfile.Project) (string, *failures.Failure) {
		spec := projectFile.Secrets.GetByName(name)
		if spec == nil {
			return "", variables.FailExpandVariable.New("secrets_expand_err_spec_undefined", name)
		}

		org, failure := organizations.FetchByURLName(projectFile.Owner)
		if failure != nil {
			return "", failure
		}

		proj, failure := projects.FetchByName(org, projectFile.Name)
		if failure != nil {
			return "", failure
		}

		kp, failure := keypairs.Fetch(secretsClient)
		if failure != nil {
			return "", failure
		}

		userSecrets, failure := fetchAll(secretsClient, org)
		if failure != nil {
			return "", failure
		}

		userSecret := findBestUserSecret(userSecrets, spec, proj)
		if userSecret == nil {
			return "", variables.FailExpandVariable.New("secrets_expand_err_not_found", name)
		}

		return decodeAndDecrypt(kp, *userSecret.Value)
	}
}

func findBestUserSecret(userSecrets []*secretsModels.UserSecret, spec *projectfile.SecretSpec, project *models.Project) *secretsModels.UserSecret {
	if project == nil {
		return nil
	}

	projectIDStr := project.ProjectID.String()

	var selectedSecret *secretsModels.UserSecret
	for _, userSecret := range userSecrets {
		secretProjectIDStr := userSecret.ProjectID.String()

		if !strings.EqualFold(*userSecret.Name, spec.Name) {
			continue
		} else if *userSecret.IsUser && secretProjectIDStr == projectIDStr {
			// best possible match regardless
			return userSecret
		} else if secretProjectIDStr != "" && secretProjectIDStr != projectIDStr {
			// this is a project secret but project id's don't match
			continue
		} else if spec.IsUser && !*userSecret.IsUser {
			// user scoped secret required
			continue
		} else if spec.IsProject && !*userSecret.IsUser && secretProjectIDStr == "" {
			// org scoped secret when project or user scope required
			continue
		}

		if selectedSecret == nil {
			// requirements met and nothing else selected yet
			selectedSecret = userSecret
			continue
		} else if *userSecret.IsUser && !*selectedSecret.IsUser {
			// prefer user scoped
			selectedSecret = userSecret
			continue
		} else if secretProjectIDStr == projectIDStr {
			// prefer project scoped
			selectedSecret = userSecret
			continue
		}

	}
	return selectedSecret
}
