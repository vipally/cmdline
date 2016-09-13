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
		line = ` ping	 /l=2 127.0.0.1 --n	 1   ip2  -i=3 ip3 -r=	 5 -w =4 /k = 6 `

		s       = []string{"", "", ""}
		n       = []int{0, 0, 0, 0, 0, 0}
		s_check = []string{"127.0.0.1", "ip2", "ip3"}
		n_check = []int{1, 2, 3, 4, 5, 6}
	)

	argv := cmdline.SplitLine(line)
	cmd := cmdline.NewFlagSet("cmdline", cmdline.PanicOnError)
	cmd.StringVar(&s[0], "", "ip", "", true, "ip")
	cmd.StringVar(&s[1], "", "ip2", "", true, "ip2")
	cmd.StringVar(&s[2], "", "ip3", "", true, "ip3")
	cmd.IntVar(&n[0], "n", "n", -1, true, "count")
	cmd.IntVar(&n[1], "l", "l", -1, true, "size")
	cmd.IntVar(&n[2], "i", "i", -1, true, "TTL")
	cmd.IntVar(&n[3], "w", "w", -1, true, "timeout")
	cmd.IntVar(&n[4], "r", "r", -1, true, "count")
	cmd.IntVar(&n[5], "k", "k", -1, true, "host-list")
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
