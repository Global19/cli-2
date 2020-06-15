package sscommon

import (
	"os/exec"
	"strings"
	"syscall"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/osutils"
)

var escaper *osutils.ShellEscape

func init() {
	escaper = osutils.NewBatchEscaper()
}

var lineBreak = "\r\n"
var lineBreakChar = `\r\n`

func stop(cmd *exec.Cmd) *failures.Failure {
	// windows should use "CTRL_CLOSE_EVENT"; SIGKILL works
	sig := syscall.SIGKILL

	// may panic if process no longer exists
	defer failures.Recover()
	if err := cmd.Process.Signal(sig); err != nil {
		return FailSignalCmd.Wrap(err)
	}

	return nil
}

// EscapeEnv escapes all values so they can be exported
func EscapeEnv(env map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range env {
		result[k] = v
		result[k] = escaper.Escape(result[k])
		result[k] = strings.ReplaceAll(result[k], lineBreak, lineBreakChar)
	}
	return result
}
