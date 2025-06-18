package main

import (
	"fmt"
	"os"

	acruncmd "github.com/takei0107/acrun/cmd"
)

func main() {
	a, err := acruncmd.ParseCmdArgs()
	if err != nil {
		if err, ok := err.(*acruncmd.InvalidArgsError); ok {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			acruncmd.PrintUsage()
			os.Exit(1)
		}

		panic(err)
	}

	p := a.ToCmdRunParam()
	if err := acruncmd.Run(p); err != nil {
		panic(err)
	}
}
