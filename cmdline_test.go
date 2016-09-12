package cmdline_test

import (
	"fmt"
	"testing"

	"github.com/vipally/cmdline"
)

func TestGetUsage(t *testing.T) {
	msg := `Usage of [cmdline.test.exe]:
  Summary:
    summary

  Usage:
    cmdline.test.exe [-b=<b>] [-s=<s>] [-test.v=<s>]
  -b=<b> (default true)	b
  -s=<s>  string	s
  -test.v=<s>  string
      s

  CopyRight:
    copyright

  Details:
    details
`
	cmdline.Summary("summary")
	cmdline.Details("details")
	cmdline.CopyRight("copyright")
	cmdline.Bool("b", "b", true, false, "b")
	cmdline.String("s", "s", "", false, "s")
	cmdline.String("test.v", "s", "", false, "s")
	cmdline.Parse()
	if cmdline.GetUsage() != msg {
		fmt.Println(cmdline.GetUsage())
		fmt.Println(msg)
		t.Errorf("GetUsage() get unexpect string")
	}
	//fmt.Println(cmdline.GetUsage())
}

func TestNonameFlag(t *testing.T) {

	var (
		line = ` ping	 /n2=2 127.0.0.1 --n1	 1   ip2  -n3=3 ip3 -n5=	 5 -n4 =4 /n6 = 6 `

		s       = []string{"", "", ""}
		n       = []int{0, 0, 0, 0, 0, 0}
		s_check = []string{"127.0.0.1", "ip2", "ip3"}
		n_check = []int{1, 2, 3, 4, 5, 6}
	)

	argv := cmdline.SplitLine(line)
	cmd := cmdline.NewFlagSet("cmdline", cmdline.PanicOnError)
	cmd.StringVar(&s[0], "", "ip", "", true, "")
	cmd.StringVar(&s[1], "", "ip2", "", true, "")
	cmd.StringVar(&s[2], "", "ip3", "", true, "")
	cmd.IntVar(&n[0], "n1", "n1", -1, true, "")
	cmd.IntVar(&n[1], "n2", "n2", -1, true, "")
	cmd.IntVar(&n[2], "n3", "n3", -1, true, "")
	cmd.IntVar(&n[3], "n4", "n4", -1, true, "")
	cmd.IntVar(&n[4], "n5", "n5", -1, true, "")
	cmd.IntVar(&n[5], "n6", "n6", -1, true, "")
	cmd.Parse(argv[1:])

	for i, v := range s {
		if v != s_check[i] {
			t.Error(i, v, s_check[i])
		}
	}
	for i, v := range n {
		if v != n_check[i] {
			t.Error(i, v, n_check[i])
		}
	}
	//fmt.Println(cmd.GetUsage())
}
