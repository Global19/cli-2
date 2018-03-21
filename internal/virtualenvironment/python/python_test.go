package python

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ActiveState/ActiveState-CLI/internal/artefact"
	"github.com/ActiveState/ActiveState-CLI/internal/config"
	"github.com/ActiveState/ActiveState-CLI/internal/distribution"
	"github.com/ActiveState/ActiveState-CLI/internal/environment"
)

func setup(t *testing.T) {
	root, _ := environment.GetRootPath()
	os.Chdir(filepath.Join(root, "test"))
}

func TestLanguage(t *testing.T) {
	venv := &VirtualEnvironment{}
	assert.Equal(t, "python", venv.Language(), "Should return python")
}

func TestDataDir(t *testing.T) {
	venv := &VirtualEnvironment{}
	assert.Empty(t, venv.DataDir())

	venv.SetDataDir("/foo")
	assert.NotEmpty(t, venv.DataDir(), "Should set the datadir")
}

func TestLanguageMeta(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}
	assert.Nil(t, venv.Artefact(), "Should not have artefact info")

	venv.SetArtefact(&artefact.Artefact{
		Meta: &artefact.Meta{
			Name: "test",
		},
		Path: "test",
	})
	assert.NotNil(t, venv.Artefact(), "Should have artefact info")
}

func TestLoadPackageFromPath(t *testing.T) {
	venv := &VirtualEnvironment{}

	datadir := filepath.Join(os.TempDir(), "as-state-test")
	os.RemoveAll(datadir)
	os.Mkdir(datadir, os.ModePerm)
	venv.SetDataDir(datadir)

	dist, fail := distribution.Obtain()
	assert.NoError(t, fail.ToError())

	var language *artefact.Artefact
	for _, lang := range dist.Languages {
		if lang.Meta.Name == venv.Language() {
			language = lang
			break
		}
	}

	assert.Nil(t, language, "Language should be nil as we don't have python artifacts yet")

	// artf := dist.Artefacts[dist.Languages[0].Hash][0]
	// fail = venv.LoadArtefact(artf)
	// assert.NoError(t, fail.ToError(), "Loads artefact without errors")

	// // Todo: Test with datadir as source, not the archived version
	// assert.FileExists(t, filepath.Join(datadir, "lib", artf.Hash, "artefact.json"), "Should create a package symlink")
}

func TestActivate(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}

	venv.SetArtefact(&artefact.Artefact{
		Meta: &artefact.Meta{
			Name:    "python",
			Version: "2.7.11",
		},
		Path: "test",
	})

	datadir := config.GetDataDir()
	datadir = filepath.Join(datadir, "test")
	venv.SetDataDir(datadir)

	venv.Activate()

	assert.DirExists(t, filepath.Join(venv.DataDir(), "bin"))
	assert.DirExists(t, filepath.Join(venv.DataDir(), "lib"))
}
