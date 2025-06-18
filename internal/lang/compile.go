package lang

import (
	"os/exec"
	"strings"

	"github.com/takei0107/acrun/internal/util"
)

type compileCmd struct {
	cmd  string
	args []string
}

func (cc *compileCmd) runCompile() (*cmdResult, error) {
	cmd := exec.Command(cc.cmd, cc.args...)

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

	go util.ReadToOuts(stdout, &result.outs)
	go util.ReadToOuts(stderr, &result.errors)

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

func (cc *compileCmd) String() string {
	s := make([]string, len(cc.args)+1, len(cc.args)+1)
	i := 0
	s[i] = cc.cmd
	i++
	for _, arg := range cc.args {
		s[i] = arg
		i++
	}
	return strings.Join(s, " ")
}
