//package cmdline support a friendly command line interface based on flag
package cmdline

import (
	"os"
	"path"
	"strings"
)

var (
	thisCmd    = get_cmd(os.Args[0])
	workDir, _ = os.Getwd()
)

func format_path(s string) string {
	ss := strings.Replace(s, "\\", "/", -1)
	return ss
}

func get_cmd(arg0 string) string {
	cmd := format_path(arg0)
	d := path.Dir(cmd) + "/"
	app := strings.TrimPrefix(cmd, d)
	return app
}

func WorkDir() string {
	return workDir
}
func Exit(code int) {
	os.Exit(code)
}

//<version>
//<buildtime>
//<thiscmd>
//replace this tag to proper string
func ReplaceTags(s string) string {
	s = strings.Replace(s, "<thiscmd>", thisCmd, -1)
	s = strings.Replace(s, "<buildtime>", "", -1)
	s = strings.Replace(s, "<version>", "", -1)
	return s
}
