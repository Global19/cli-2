package bash

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/pkg/projectfile"
)

// SubShell covers the subshell.SubShell interface, reference that for documentation
type SubShell struct {
	binary string
	rcFile *os.File
	cmd    *exec.Cmd
	wg     *sync.WaitGroup
}

// Shell - see subshell.SubShell
func (v *SubShell) Shell() string {
	return "bash"
}

// Binary - see subshell.SubShell
func (v *SubShell) Binary() string {
	return v.binary
}

// SetBinary - see subshell.SubShell
func (v *SubShell) SetBinary(binary string) {
	v.binary = binary
}

// RcFile - see subshell.SubShell
func (v *SubShell) RcFile() *os.File {
	return v.rcFile
}

// SetRcFile - see subshell.SubShell
func (v *SubShell) SetRcFile(rcFile *os.File) {
	v.rcFile = rcFile
}

// RcFileExt - see subshell.SubShell
func (v *SubShell) RcFileExt() string {
	return ""
}

// RcFileTemplate - see subshell.SubShell
func (v *SubShell) RcFileTemplate() string {
	return "bashrc.sh"
}

// Activate - see subshell.SubShell
func (v *SubShell) Activate(wg *sync.WaitGroup) error {
	v.wg = wg
	wg.Add(1)

	shellArgs := []string{"--rcfile", v.rcFile.Name()}
	cmd := exec.Command(v.Binary(), shellArgs...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Start()

	v.cmd = cmd

	var err error
	go func() {
		err = cmd.Wait()
		v.wg.Done()
	}()

	return err
}

// Deactivate - see subshell.SubShell
func (v *SubShell) Deactivate() error {
	if !v.IsActive() {
		return nil
	}

	var err error
	func() {
		// Go's Process.Kill is not very safe to use, it throws a panic if the process no longer exists
		defer failures.Recover()
		err = v.cmd.Process.Kill()
	}()

	if err == nil {
		v.cmd = nil
	}
	return err
}

// Run - see subshell.SubShell
func (v *SubShell) Run(script string) error {
	tmpfile, err := ioutil.TempFile("", "bash-script")
	if err != nil {
		return err
	}

	tmpfile.WriteString("#!/usr/bin/env bash\n")
	tmpfile.WriteString(script)
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())
	os.Chmod(tmpfile.Name(), 0755)

	runCmd := exec.Command(tmpfile.Name())
	runCmd.Dir = filepath.Dir(projectfile.Get().Path())
	runCmd.Stdin, runCmd.Stdout, runCmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	return runCmd.Run()
}

// IsActive - see subshell.SubShell
func (v *SubShell) IsActive() bool {
	return v.cmd != nil && (v.cmd.ProcessState == nil || !v.cmd.ProcessState.Exited())
}
