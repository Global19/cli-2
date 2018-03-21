package python

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ActiveState/ActiveState-CLI/internal/artefact"
	"github.com/ActiveState/ActiveState-CLI/internal/failures"
	"github.com/ActiveState/ActiveState-CLI/internal/fileutils"
)

// VirtualEnvironment covers the virtualenvironment.VirtualEnvironment interface, reference that for documentation
type VirtualEnvironment struct {
	datadir  string
	artefact *artefact.Artefact
}

// Language - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) Language() string {
	return "python"
}

// DataDir - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) DataDir() string {
	return v.datadir
}

// SetDataDir - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) SetDataDir(path string) {
	v.datadir = path
}

// Artefact - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) Artefact() *artefact.Artefact {
	return v.artefact
}

// SetArtefact - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) SetArtefact(artf *artefact.Artefact) {
	v.artefact = artf
}

// LoadArtefact - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) LoadArtefact(artf *artefact.Artefact) *failures.Failure {
	switch artf.Meta.Type {
	case "package":
		return v.loadPackage(artf)
	default:
		return failures.FailUser.New("err_language_not_supported", artf.Meta.Name)
	}
}

func (v *VirtualEnvironment) loadPackage(artf *artefact.Artefact) *failures.Failure {
	if err := fileutils.Mkdir(v.datadir, "lib"); err != nil {
		return failures.FailIO.Wrap(err)
	}

	artfPath := filepath.Dir(artf.Path)
	err := filepath.Walk(artfPath, func(subpath string, f os.FileInfo, err error) error {
		subpath = strings.TrimPrefix(subpath, artfPath)
		if subpath == "" {
			return nil
		}
		target := filepath.Join(v.DataDir(), "lib", filepath.Base(artfPath), subpath)
		if err := fileutils.Mkdir(filepath.Dir(target), "lib"); err != nil {
			return failures.FailIO.Wrap(err)
		}

		return os.Symlink(filepath.Join(artfPath, subpath), target)
	})

	if err != nil {
		return failures.FailIO.Wrap(err)
	}

	return nil
}

// Activate - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) Activate() *failures.Failure {
	if err := fileutils.Mkdir(v.datadir, "bin"); err != nil {
		return err
	}
	return fileutils.Mkdir(v.datadir, "lib")
}

// Env - see virtualenvironment.VirtualEnvironment
func (v *VirtualEnvironment) Env() map[string]string {
	path := filepath.Join(v.datadir, "language", "bin") + string(os.PathListSeparator) + os.Getenv("PATH")
	path = filepath.Join(v.datadir, "bin") + string(os.PathListSeparator) + path
	return map[string]string{
		"PYTHONPATH": filepath.Join(v.datadir, "lib"),
		"PATH":       path,
	}
}
