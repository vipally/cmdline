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
	version    = "0.0.1"
)

func format_path(s string) string {
	return path.Clean(s)
}

func get_cmd(arg0 string) string {
	cmd := format_path(arg0)
	d := path.Dir(cmd) + "/"
	app := strings.TrimPrefix(cmd, d)
	app = strings.TrimSuffix(app, ".exe")
	return app
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
