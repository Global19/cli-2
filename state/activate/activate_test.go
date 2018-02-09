package activate

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	assert := assert.New(t)

	Cc := Command.GetCobraCmd()
	Cc.SetArgs([]string{"activate"})

	Command.Execute()

	assert.Equal(true, true, "Execute didn't panic")
}

func TestExecuteGitClone(t *testing.T) {
	cwd, _ := os.Getwd() // store
	repo, err := filepath.Abs(filepath.Join("..", "..", "internal", "scm", "git", "testdata", "repo"))
	assert.Nil(t, err, "The test Git repository exists")

	tempdir, err := ioutil.TempDir("", "ActiveState-CLI-")
	assert.Nil(t, err, "A temporary directory was created")
	err = os.Chdir(tempdir)
	assert.Nil(t, err, "Changed into temporary directory")

	// Test basic clone.
	_, err = os.Stat("repo")
	Flags.Path = ""
	assert.True(t, os.IsNotExist(err), "The cloned repository does not exist yet")
	Command.GetCobraCmd().SetArgs([]string{repo})
	Command.Execute()
	_, err = os.Stat(filepath.Join(tempdir, "repo"))
	assert.Nil(t, err, "The cloned repository exists")
	files := []string{"foo.txt", "bar.txt", "baz.txt"}
	for _, file := range files {
		_, err = os.Stat(filepath.Join(tempdir, "repo", file))
		assert.Nil(t, err, "The cloned repository contains an expected file")
	}

	// Test clone with specified directory.
	_, err = os.Stat("repo2")
	Flags.Path = ""
	os.Chdir(tempdir)
	assert.True(t, os.IsNotExist(err), "The cloned repository does not exist yet")
	Command.GetCobraCmd().SetArgs([]string{repo, "--path", "repo2"})
	Command.Execute()
	newCwd, _ := os.Getwd()
	assert.Equal(t, "repo2", filepath.Base(newCwd), "The cloned repository exists and was changed into")
	_, err = os.Stat(filepath.Join(tempdir, "repo2"))
	assert.Nil(t, err, "The cloned repository exists")
	for _, file := range files {
		_, err = os.Stat(filepath.Join(tempdir, "repo2", file))
		assert.Nil(t, err, "The cloned repository contains an expected file")
	}

	// Test clone of invalid repository.
	Flags.Path = ""
	os.Chdir(tempdir)
	Command.GetCobraCmd().SetArgs([]string{cwd})
	Command.Execute()
	// TODO: assert that an error occurred

	err = os.Chdir(cwd) // restore
	assert.Nil(t, err, "Changed back to original directory")
	err = os.RemoveAll(tempdir) // clean up
	assert.Nil(t, err, "The temporary directory was removed")
}

func TestExecuteGitCloneRemote(t *testing.T) {
	cwd, _ := os.Getwd() // store

	tempdir, err := ioutil.TempDir("", "ActiveState-CLI-")
	assert.Nil(t, err, "A temporary directory was created")
	err = os.Chdir(tempdir)
	assert.Nil(t, err, "Changed into temporary directory")

	Flags.Path = ""
	os.Chdir(tempdir)
	Command.GetCobraCmd().SetArgs([]string{"https://github.com/ActiveState/repo"})
	Command.Execute()
	_, err = os.Stat(filepath.Join(tempdir, "repo"))
	assert.Nil(t, err, "The cloned repository exists")
	files := []string{"foo.txt", "bar.txt", "baz.txt"}
	for _, file := range files {
		_, err = os.Stat(file)
		assert.Nil(t, err, "The cloned repository contains an expected file")
	}

	Flags.Path = ""
	os.Chdir(tempdir)
	Command.GetCobraCmd().SetArgs([]string{"https://github.com/ActiveState/does-not-exist", "--path", "repo2"})
	Command.Execute()
	_, err = os.Stat("repo2")
	assert.Error(t, err, "The non-existant repository did not have an ActiveState config file; no clone happened")

	err = os.Chdir(cwd) // restore
	assert.Nil(t, err, "Changed back to original directory")
	err = os.RemoveAll(tempdir) // clean up
	assert.Nil(t, err, "The temporary directory was removed")
}
