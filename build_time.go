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

//time that this package build
func BuildTime() string {
	return buildTime
}
