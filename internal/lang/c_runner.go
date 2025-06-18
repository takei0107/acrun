package lang

type cRunner struct {
	abstractRunner
}

func newCRunner(rc *RunnerConfig) (Runner, error) {
	r := new(cRunner)
	r.isNeedCompile = rc.IsNeedCompile
	r.compileCmd = &compileCmd{
		cmd:  rc.CompileCmd.Cmd,
		args: rc.CompileCmd.Args,
	}
	r.exeCmd = &exeCmd{
		cmd:  rc.ExeCmd.Cmd,
		args: rc.ExeCmd.Args,
	}

	return r, nil
}
