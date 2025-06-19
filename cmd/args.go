package acruncmd

import (
	"flag"
	"fmt"
)

type InvalidArgsError struct {
	argCount int
}

func (err *InvalidArgsError) Error() string {
	return fmt.Sprintf("2 arguments is required. but got=%d\n", err.argCount)
}

type options struct {
	contest      string
	questionTask string
	fileName     string
	exe          string
}

type arguments struct {
	questionId string
	lang       string
}

type parsedArgs struct {
	opts *options
	args *arguments
}

func (a *parsedArgs) ToCmdRunParam() *cmdRunParam {
	p := new(cmdRunParam)

	p.contest = a.opts.contest
	p.questionId = a.args.questionId
	p.lang = a.args.lang
	p.fileName = a.opts.fileName
	p.exe = a.opts.exe

	return p
}

func parseOptions() *options {
	flag.Usage = Usage

	c := flag.String("c", "", "contest name (default: current dir name)")
	t := flag.String("t", "", "contest question task (default: value of question) ")
	f := flag.String("f", "", "file name")
	e := flag.String("e", "", "exec comand name")

	flag.Parse()

	opts := new(options)

	opts.contest = *c
	opts.questionTask = *t
	opts.fileName = *f
	opts.exe = *e

	return opts
}

func parseArgs() (*arguments, error) {
	args := flag.Args()
	if len(args) < 2 {
		err := new(InvalidArgsError)
		err.argCount = len(args)
		return nil, err
	}

	ap := new(arguments)
	ap.lang = args[0]
	ap.questionId = args[1]

	return ap, nil
}

func ParseCmdArgs() (*parsedArgs, error) {
	opts := parseOptions()

	args, err := parseArgs()
	if err != nil {
		return nil, err
	}

	p := new(parsedArgs)
	p.opts = opts
	p.args = args

	return p, nil
}

func Usage() {
	fmt.Printf("%s\n", "usage: acrun [...options] lang question")
	flag.PrintDefaults()
}
