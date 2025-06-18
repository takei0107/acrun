package acruncmd

import (
	"fmt"

	"github.com/takei0107/acrun/internal/contest"
	"github.com/takei0107/acrun/internal/lang"
	"github.com/takei0107/acrun/internal/util"
)

type cmdRunParam struct {
	lang         string
	fileName     string
	exe          string
	contest      string
	questionTask string
	questionId   string
}

func (p *cmdRunParam) toCmdRunParamWithFt() *cmdRunParamWithFt {
	r := new(cmdRunParamWithFt)

	r.cmdRunParam = p
	r.ft = lang.ResolveFileType(p.lang)

	return r
}

func (p *cmdRunParam) toMergedParam() (*mergedParam, error) {
	pp := p.toCmdRunParamWithFt()

	jp, err := lang.GetJsonParam(pp.ft)
	if err != nil {
		return nil, err
	}

	mp := pp.mergeWith(jp)
	mp.replaceParams()

	return mp, nil
}

type cmdRunParamWithFt struct {
	*cmdRunParam
	ft lang.FileType
}

func (p *cmdRunParamWithFt) mergeWith(jp *lang.JsonParam) *mergedParam {
	mp := new(mergedParam)

	mp.ft = p.ft

	fn := p.fileName
	if fn == "" {
		fn = jp.DefaultFileName
	}
	mp.fileName = fn

	e := p.exe
	if e == "" {
		e = jp.DefaultExe
	}
	mp.exe = e

	mp.isNeedCompile = jp.IsNeedCompile

	cp := new(compileParam)
	cp.cmd = jp.Compile.Cmd
	cp.args = jp.Compile.Args
	mp.compileParam = cp

	rp := new(runParam)
	rp.cmd = jp.Run.Cmd
	rp.args = jp.Run.Args
	mp.runParam = rp

	mp.contest = p.contest
	mp.questionTask = p.questionTask
	mp.questionId = p.questionId

	return mp
}

type replacePatternMap = map[string]string

type compileParam struct {
	cmd  string
	args []string
}

func (c *compileParam) replacePatterns(p replacePatternMap) {
	r := util.NewReplacer()

	for pp, v := range p {
		r.AddReplacements(pp, v)
	}

	c.cmd = r.ReplaceStr(c.cmd)
	c.args = r.ReplaceStrSlice(c.args)
}

type runParam struct {
	cmd  string
	args []string
}

func (rp *runParam) replacePatterns(p replacePatternMap) {
	r := util.NewReplacer()

	for pp, v := range p {
		r.AddReplacements(pp, v)
	}

	rp.cmd = r.ReplaceStr(rp.cmd)
	rp.args = r.ReplaceStrSlice(rp.args)
}

type mergedParam struct {
	ft            lang.FileType
	fileName      string
	exe           string
	isNeedCompile bool
	compileParam  *compileParam
	runParam      *runParam
	contest       string
	questionTask  string
	questionId    string
}

func (mp *mergedParam) replaceParams() {
	rp := map[string]string{
		"%fileName%": mp.fileName,
		"%exe%":     mp.exe,
	}
	mp.compileParam.replacePatterns(rp)
	mp.runParam.replacePatterns(rp)
}

func (mp *mergedParam) toRunnerConfig() *lang.RunnerConfig {
	c := new(lang.RunnerConfig)
	c.Ft = mp.ft
	c.FileName = mp.fileName

	c.IsNeedCompile = mp.isNeedCompile

	cc := new(lang.CompileCmd)
	cc.Cmd = mp.compileParam.cmd
	cc.Args = mp.compileParam.args
	c.CompileCmd = cc

	ec := new(lang.ExeCmd)
	ec.Cmd = mp.runParam.cmd
	ec.Args = mp.runParam.args
	c.ExeCmd = ec

	return c
}

func (mp *mergedParam) toQuestionConfig() (*contest.QuestionConfig, error) {

	cn, err := contest.ResolveContestName(mp.contest)
	if err != nil {
		return nil, err
	}

	c := new(contest.QuestionConfig)

	c.Contest = cn
	c.QuestionId = mp.questionId
	if mp.questionTask == "" {
		c.QuestionTask = cn
	} else {
		c.QuestionTask = mp.questionTask
	}

	// FIXME: リソースタイプの汎用性を持たせたい
	c.ResourceType = contest.ResourceOfHtml

	return c, nil
}

func compile(runner lang.Runner) error {
	fmt.Printf("[acrun] start compile\n")

	err := runner.Compile()
	if err != nil {
		return err
	}

	fmt.Printf("[acrun] end compile\n")

	return nil
}

func run(runner lang.Runner, s []*contest.SampleInOuts, qc *contest.QuestionConfig) error {
	fmt.Printf("[acrun] start run\n")

	fmt.Printf("[acrun]   contest=%s, task=%s, id=%s\n", qc.Contest, qc.QuestionTask, qc.QuestionId)
	for _, ss := range s {
		fmt.Printf("[acrun]   sample-no=%d\n", ss.No)
		err := runner.Run(ss.Inputs, ss.Outputs)
		if err != nil {
			return err
		}
	}

	fmt.Printf("[acrun] end run\n")

	return nil
}

func Run(p *cmdRunParam) error {

	mp, err := p.toMergedParam()
	if err != nil {
		return err
	}

	rc := mp.toRunnerConfig()

	runner, err := lang.GetRunner(rc)
	if err != nil {
		return err
	}

	err = compile(runner)
	if err != nil {
		return err
	}

	qc, err := mp.toQuestionConfig()
	if err != nil {
		return err
	}

	s, err := contest.GetSampleInOutsSlice(qc)
	if err != nil {
		return err
	}

	err = run(runner, s, qc)
	if err != nil {
		return err
	}

	return nil
}
