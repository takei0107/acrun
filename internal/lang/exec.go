package lang

import (
	"io"
	"os/exec"

	"github.com/takei0107/acrun/internal/util"
)

type exeCmd struct {
	cmd  string
	args []string
}

func (ec *exeCmd) runExec(i []string) (*cmdResult, error) {
	cmd := exec.Command(ec.cmd, ec.args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	result := new(cmdResult)
	result.exitCode = 0

	go func() {
		defer stdin.Close()
		for _, ii := range i {
			io.WriteString(stdin, ii)
		}
	}()

	go util.ReadToOuts(stdout, result.outs)
	go util.ReadToOuts(stderr, result.errors)

	err = cmd.Wait()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			return nil, err
		}
		result.exitCode = exitErr.ExitCode()
	}

	return result, nil
}
