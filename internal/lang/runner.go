package lang

import (
	"fmt"
	"os"
)

type CompileCmd struct {
	Cmd  string
	Args []string
}

type ExeCmd struct {
	Cmd  string
	Args []string
}

type RunnerConfig struct {
	Ft            FileType
	FileName      string
	IsNeedCompile bool
	CompileCmd    *CompileCmd
	ExeCmd        *ExeCmd
}

type Runner interface {
	Compile() error
	Run(i []string, o []string) error
}

type cmdResult struct {
	exitCode int
	outs     []string
	errors   []string
}

type abstractRunner struct {
	isNeedCompile bool
	compileCmd    *compileCmd
	exeCmd        *exeCmd
}

func (r *abstractRunner) Compile() error {
	if !r.isNeedCompile {
		fmt.Printf("[acrun]   skip compile\n")
		return nil
	}

	fmt.Printf("[acrun]   compile command=\"%s\"\n", r.compileCmd)

	result, err := r.compileCmd.runCompile()
	if err != nil {
		return err
	}

	var rt error = nil
	if result.exitCode > 0 {
		fmt.Fprintf(os.Stderr, "[acrun]   command failed with code=%d\n", result.exitCode)
		fmt.Fprintf(os.Stderr, "[acrun]   errors:\n")
		for _, line := range result.errors {
			fmt.Fprintf(os.Stderr, "[acrun]     %s\n", line)
		}
		rt = fmt.Errorf("comand failed")
	}

	fmt.Printf("[acrun]   stdouts:\n")
	for _, line := range result.outs {
		fmt.Printf("[acrun]     %s\n", line)
	}

	return rt
}

func (r *abstractRunner) Run(i []string, o []string) error {
	fmt.Printf("[acrun]   run command=\"%s\"\n", r.exeCmd.cmd)

	result, err := r.exeCmd.runExec(i)
	if err != nil {
		return err
	}

	if result.exitCode > 0 {
		fmt.Fprintf(os.Stderr, "[acrun]   command failed with code=%d\n", result.exitCode)
		fmt.Fprintf(os.Stderr, "[acrun]   errors:\n")
		for _, line := range result.errors {
			fmt.Fprintf(os.Stderr, "[acrun]     %s\n", line)
		}
		return fmt.Errorf("comand failed")
	}

	fmt.Printf("[acrun]   run command succeess\n")
	fmt.Printf("[acrun]   test...\n")

	ok := true
	for j, got := range result.outs {
		if got != o[j] {
			ok = false
		}
	}

	if !ok {
		fmt.Printf("[acrun]   failed! please check outputs.\n")
	} else {
		fmt.Printf("[acrun]   success!\n")
	}

	fmt.Printf("[acrun]   sample outputs:\n")
	for _, oo := range o {
		fmt.Printf("%s\n", oo)
	}

	fmt.Printf("[acrun]   cmd outputs:\n")
	for _, oo := range result.outs {
		fmt.Printf("%s\n", oo)
	}

	return nil
}

func GetRunner(c *RunnerConfig) (Runner, error) {
	var (
		runner Runner
		err    error
	)
	switch c.Ft {
	case C:
		runner, err = newCRunner(c)
	default:
		runner = nil
		err = fmt.Errorf("command runner has not got. fileType=%v\n", c.Ft)
	}
	return runner, err
}
