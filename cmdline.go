/*
	package cmdline extends std.flag to support more feature.

	main change list from std.flag:
	1. Add LogicName and Required field for every flag, and modify the flag define interface
	2. Add Summary and Details and Version infor commond line info
	3. Add <thiscmd> <build> <version> labels for Summary and Details to get runtime info
	4. Add interface GetUsage() string
	5. Modify the Parse() logic
	6. Add noname-flag support
	7. Add /flag support
	8. Fix "-flag = x" or "-flag= x" or "-flag =x" cause panic bug
	9. Add synonyms support for with-name flags

	Usage as follow:

	cmdline.Version("1.0.2")
	cmdline.Summary("<thiscmd> is an example of cmdline package usage.")
	cmdline.Details(`Version   :<version>
    BulidTime :<buildtime>
    <thiscmd> is an example usage of github.com/vipally/cmdline package.`)
	cmdline.CopyRight("no copyright defined")

	//noname flag and require ones
	cmdline.StringVar(&target_name, "", "target_name", "", true, "target host ip or name")

	cmdline.BoolVar(&v4, "4", "v4", v4, false, "ipv4")

	//synonym with the same variables
	cmdline.IntVar(&ttl, "t", "ttl", ttl, false, "ttl")
	cmdline.IntVar(&ttl, "ttl", "synonym of -t", ttl, true, "this is synonym of -t")

	//define a synonym with method AnotherName
	c := cmdline.Int("c", "count", 0, false, "count")
	cmdline.AnotherName("count", "c")

	cmdline.Parse()

	fmt.Println(target_name, v4, ttl, *c)
	fmt.Println(cmdline.GetUsage())

	//output:
	//Usage of ([ping] Build [Sep 29 2016 21:14:37]):
	//  Summary:
	//    ping is an example of cmdline package usage.

	//  Usage:
	//    ping [-4=<v4>] [-c|count=<count>] [-t|ttl=<ttl>] <target_name>
	//  -4=<v4>	ipv4
	//  -c|count=<count>  int
	//      count
	//  -t|ttl=<ttl>  int
	//      ttl
	//  <target_name>  required  string
	//      target host ip or name

	//  CopyRight:
	//    no copyright defined

	//  Details:
	//    Version   :1.0.2
	//    BulidTime :[Sep 29 2016 21:14:37]
	//    ping is an example usage of github.com/vipally/cmdline package.
*/
package cmdline

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	thisCmd    = get_cmd(os.Args[0])
	workDir, _ = os.Getwd()
	version    = "unknown"
)

func format_path(s string) string {
	return path.Clean(s)
}

func get_cmd(arg0 string) string {
	ext := filepath.Ext(arg0)
	return strings.TrimSuffix(filepath.Base(arg0), ext)
}

func WorkDir() string {
	return workDir
}

func Exit(code int) {
	os.Exit(code)
}

func Version(v string) (old string) {
	old, version = version, v
	return
}

//<version>
//<buildtime>
//<thiscmd>
//replace this tag to proper string
func ReplaceTags(s string) string {
	s = strings.Replace(s, "<thiscmd>", thisCmd, -1)
	s = strings.Replace(s, "<buildtime>", BuildTime(), -1)
	s = strings.Replace(s, "<version>", version, -1)
	return s
}
