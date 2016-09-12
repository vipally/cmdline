package cmdline_test

import (
	"fmt"
	"testing"

	"github.com/vipally/cmdline"
)

func TestSplitLine(t *testing.T) {
	s := `ping 127.0.0.1 		 -n 	2   " --x = 5 "  	a`
	cmd := cmdline.SplitLine(s)
	result := []string{
		"ping",
		"127.0.0.1",
		"-n",
		"2",
		" --x = 5 ",
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
	fmt.Println(len(cmd), len(result))
	t.Error("SplitLine fail")
}
