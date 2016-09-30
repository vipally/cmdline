package cmdline

/*
const char* build_time(void)
{
    static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
	return psz_build_time;
}
*/
import "C"

var (
	buildTime = C.GoString(C.build_time())
)

//BuildTime returns the package build time from cgo, eg [Sep 30 2016 00:50:14]
func BuildTime() string {
	return buildTime
}
