package preprocess

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ActiveState/cli/internal/constants"
)

// Constants holds constants that will be preprocessed, meaning the key value parts here will be built into the constants
// package as actual constants, using the build-time interpretations
var Constants = map[string]func() string{}

func init() {
	branchName, branchNameFull := branchName()
	buildNumber := buildNumber()

	incrementer, err := NewVersionIncrementer(NewGithubProvider(os.Getenv("GITHUB_REPO_TOKEN")), branchName, buildEnvironment())
	if err != nil {
		log.Fatalf("Could not initialize version incrementer: %s", err)
	}

	Constants["BranchName"] = func() string { return branchName }
	Constants["BuildNumber"] = func() string { return buildNumber }
	Constants["RevisionHash"] = func() string { return getCmdOutput("git rev-parse --verify " + branchNameFull) }
	Constants["RevisionHashShort"] = func() string { return getCmdOutput("git rev-parse --short " + branchNameFull) }
	Constants["Version"] = func() string { return mustIncrementVersionRevision(incrementer, Constants["RevisionHashShort"]()) }
	Constants["VersionNumber"] = func() string { return mustIncrementVersion(incrementer) }
	Constants["IncrementString"] = func() string { return mustGetIncrementString(incrementer) }
	Constants["Date"] = func() string { return time.Now().Format("Mon Jan 2 2006 15:04:05 -0700 MST") }
	Constants["UserAgent"] = func() string {
		return fmt.Sprintf("%s/%s; %s", constants.CommandName, Constants["Version"](), branchName)
	}
	Constants["APITokenName"] = func() string { return fmt.Sprintf("%s-%s", constants.APITokenNamePrefix, branchName) }
}

// gitBranchName returns the branch name of the current git commit / PR
func gitBranchName() string {
	// branch name variable set by Azure CI during pull request
	if branch, isset := os.LookupEnv("SYSTEM_PULLREQUEST_SOURCEBRANCH"); isset {
		return "origin/" + branch
	}
	// branch name variable set by Azure CI
	if branch, isset := os.LookupEnv("BUILD_SOURCEBRANCH"); isset {
		return "origin/" + strings.TrimPrefix(branch, "refs/heads/")
	}
	branch := getCmdOutput("git rev-parse --abbrev-ref HEAD")
	return branch
}

// branchName returns the release name and the branch name it is generated from
// Usually the release name is identical to the branch name, unless environment variable
// `BRANCH_OVERRIDE` is set
func branchName() (string, string) {
	branch := gitBranchName()
	releaseName := strings.TrimPrefix(branch, "origin/")

	if releaseOverride, isset := os.LookupEnv("BRANCH_OVERRIDE"); isset {
		releaseName = releaseOverride
	}

	return releaseName, branch
}

func buildNumber() string {
	out := getCmdOutput("git rev-list --all --count")
	return strings.TrimSpace(out)
}

func getCmdOutput(cmdString string) string {
	cmdArgs := strings.Split(cmdString, " ")

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed, command: %s, args: %v, output: %s, error: %s, code: %s", cmdArgs[0], cmdArgs[1:], out.String(), stderr.String(), err)
		os.Exit(1)
	}
	return strings.Trim(out.String(), "\n")
}

func buildEnvironment() int {
	if !onCI() {
		return localEnv
	}

	return remoteEnv
}

func onCI() bool {
	if os.Getenv("CI") != "" {
		return true
	}
	return false
}

func mustIncrementVersion(incrementer *VersionIncrementer) string {
	version, err := incrementer.IncrementVersion()
	if err != nil {
		log.Fatalf("Failed to increment version: %s", err)
	}

	return version.String()
}

func mustIncrementVersionRevision(incrementer *VersionIncrementer, revision string) string {
	version, err := incrementer.IncrementVersionRevision(revision)
	if err != nil {
		log.Fatalf("Failed to increment version: %s", err)
	}

	return version.String()
}

func mustGetIncrementString(incrementer *VersionIncrementer) string {
	increment, err := incrementer.IncrementType()
	if err != nil {
		log.Fatalf("Failed to get increment string: %s", err)
	}

	return increment
}
