//package cmdline support a friendly command line interface based on flag
package cmdline

import (
	"path"
	"strings"
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
