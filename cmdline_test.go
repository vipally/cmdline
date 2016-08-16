package cmdline_test

import (
	//"fmt"
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

  Details:
    details
`
	cmdline.Summary("summary")
	cmdline.Details("details")
	cmdline.Bool("b", "b", true, false, "b")
	cmdline.String("s", "s", "", false, "s")
	cmdline.String("test.v", "s", "", false, "s")
	cmdline.Parse()
	if cmdline.GetUsage() != msg {
		t.Errorf("GetUsage() get unexpect string")
	}
	//fmt.Println(cmdline.GetUsage())
}
