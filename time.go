package cmdline

var buildTime = "[2018-09-01 20:33:25]"

//const version = "1.0.0"

func BuildTime() string {
	return buildTime
}

func SetBuildTime(buildtime string) string {
	old := buildTime
	buildTime = buildtime
	return old
}

//func Version() string { return _version }
