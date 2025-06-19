package acruncmd

import (
	"reflect"
	"testing"
)

func Test_replacePatterns(t *testing.T) {
	tests := []struct {
		name      string
		cmd       string
		args      []string
		rpMap     replacePatternMap
		cmdWant   string
		argsWants []string
	}{
		{
			name: "no replacement pattern",
			cmd: "hoge",
			args: []string{"foo", "bar", "fizz", "buzz"},
			rpMap: replacePatternMap{
				"abc": "cdf",
				"ok":  "ng",
				"zzz": "fff",
			},
			cmdWant: "hoge",
			argsWants: []string{"foo", "bar", "fizz", "buzz"},
		},
		{
			name: "normal replacement pattern1",
			cmd: "%cmd%",
			args: []string{"%args1%", "%args2%", "%args3%"},
			rpMap: replacePatternMap{
				"%cmd%": "hoge",
				"%args1%": "foo",
				"%args2%": "bar",
				"%args3%": "buzz",
			},
			cmdWant: "hoge",
			argsWants: []string{"foo", "bar", "buzz"},
		},
		{
			name: "normal replacement pattern2",
			cmd: "a%cmd%e",
			args: []string{"12%args1%", "%args2%zzz", "987%args3%abc"},
			rpMap: replacePatternMap{
				"%cmd%": "hoge",
				"%args1%": "foo",
				"%args2%": "bar",
				"%args3%": "buzz",
			},
			cmdWant: "ahogee",
			argsWants: []string{"12foo", "barzzz", "987buzzabc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := &runParam{
				cmd:  tt.cmd,
				args: tt.args,
			}

			replacePatterns(rp, tt.rpMap)

			{
				want := tt.cmdWant
				got := rp.cmd

				if got != want {
					t.Errorf("got=%s, but want=%s\n", got, want)
				}
			}

			{
				want := tt.argsWants
				got := rp.args

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got=%v, but want=%v\n", got, want)
				}
			}
		})
	}
}
