package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/vipally/cmdline"
)

const (
	gsExpTime = "(?sm:\\\"\\[(?P<T>.*?)\\]\\\")"
)

var (
	gGoPath  = ""
	file     = "github.com/vipally/cmdline/time.go"
	gExpTime = regexp.MustCompile(gsExpTime)
	src      = []byte("$T")
)

func formatPath(path string) string {
	return filepath.ToSlash(filepath.Clean(expadGoPath(path)))
}
func expadGoPath(path string) (r string) {
	r = path
	if filepath.VolumeName(path) == "" {
		r = filepath.Join(gGoPath, path)
	}
	return
}

func main() {
	cmdline.StringVar(&file, "f", "file", file, false, "file")
	cmdline.Parse()
	s := os.Getenv("GOPATH")
	if ss := strings.Split(s, ";"); ss != nil && len(ss) > 0 {
		gGoPath = formatPath(ss[0]) + "/src/"
	}
	file = formatPath(file)

	t := time.Now()
	now := []byte(fmt.Sprintf(`"[%04d-%02d-%02d %02d:%02d:%02d]"`, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))

	if content, err := ioutil.ReadFile(file); err == nil {
		dst := gExpTime.ReplaceAll(content, now)
		//fmt.Println(string(dst))
		if f, err := os.Create(file); err == nil {
			f.Write(dst)
			f.Close()
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
