package main

import (
	acruncmd "github.com/takei0107/acrun/cmd"
)

func main() {
	a, err := acruncmd.ParseCmdArgs()
	if err != nil {
		panic(err)
	}

	p := a.ToCmdRunParam()
	if err := acruncmd.Run(p); err != nil {
		panic(err)
	}
}
