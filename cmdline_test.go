// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally

package cmdline_test

import (
	"strings"
	"testing"

	"github.com/vipally/cmdline"
)

func TestGetUsage(t *testing.T) {
	var sCheck = `Usage of ([<thiscmd>] Build [<versiontime>]):
  Summary:
    <thiscmd> is an example of cmdline package usage.

  Usage:
    <thiscmd> [-4=<v4>] [-c|count=<count>] [-t|ttl=<ttl>] <host> [<host2>]
  -4=<v4>
    ipv4
  -c|count=<count>  int
    count
  -t|ttl=<ttl>  int
    ttl
  <host>  required  string
    host ip or name
  <host2>  string
    second host ip or name

  CopyRight:
    no copyright defined

  Details:
    AppName     : <appname>
    Version     : <version>
    VersionTime : <versiontime>
    VersionTag  : <versiontag>
    Validity    : <validity>
    <thiscmd> is an example usage of github.com/vipally/cmdline package.
`

	var (
		host, host2 string
		v4          = false
		ttl         = 0
		c           = 0
	)
	cmdline.AppName("example")
	cmdline.Version("1.0.2")
	cmdline.VersionTime("2018-09-01")
	cmdline.VersionTag("abcdef0123456789")
	cmdline.Validity("2019-09-01 00:00:00")
	sCheck = cmdline.ReplaceTags(sCheck)
	cmdline.Summary("<thiscmd> is an example of cmdline package usage.")
	cmdline.Details(`AppName     : <appname>
Version     : <version>
VersionTime : <versiontime>
VersionTag  : <versiontag>
Validity    : <validity>
<thiscmd> is an example usage of github.com/vipally/cmdline package.`)
	cmdline.CopyRight("no copyright defined")

	//no-name flag and required ones
	cmdline.StringVar(&host, "", "host", "", true, "host ip or name")
	cmdline.StringVar(&host2, "", "host2", "", false, "second host ip or name")

	cmdline.BoolVar(&v4, "4", "v4", v4, false, "ipv4")

	//synonym with the same variables
	cmdline.IntVar(&ttl, "t", "ttl", ttl, false, "ttl")
	cmdline.IntVar(&ttl, "ttl", "synonym of -t", ttl, true, "this is synonym of -t")

	//define a synonym with method AnotherName
	cmdline.IntVar(&c, "c", "count", 0, false, "count")
	cmdline.Alias("count", "c")

	//cmdline.Parse()
	usage := cmdline.GetUsage()
	//fmt.Println(usage)
	if !strings.HasSuffix(usage, sCheck) {
		t.Errorf("GetUsage fail \nneed:\n%#v\ngot:\n%#v", sCheck, usage)
	}
}

func TestNonameFlag(t *testing.T) {

	var (
		line = ` ping	 /l=2 127.0.0.1 --n	 1   ip2  -i=3 ip3 -r=	 5 -w =4 /k = 666 -k2 = 6`

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
	cmd.Alias("k2", "k")
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
