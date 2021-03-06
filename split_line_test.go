// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally

package cmdline_test

import (
	"fmt"
	"testing"

	"github.com/vipally/cmdline"
)

func TestSplitLine(t *testing.T) {
	s := `ping 127.0.0.1 		 -n= 	2   " --x = 5 "  "--help"	a`
	cmd := cmdline.SplitLine(s)
	result := []string{
		"ping",
		"127.0.0.1",
		"-n=",
		"2",
		`" --x = 5 "`,
		`"--help"`,
		"a",
	}
	suc := true
	if len(cmd) == len(result) {
		for i, v := range cmd {
			if v != result[i] {
				fmt.Println(v, result[i])
				suc = false
				break
			}
		}
		if suc {
			return
		}
	}
	//fmt.Println(len(cmd), len(result))
	t.Error("SplitLine fail")
}
