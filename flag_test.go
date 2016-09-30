package cmdline_test

/*    //comment this line to open the case
import (
	"flag"
	"testing"

	"github.com/vipally/cmdline"
)

//basic requirement test case
func TestFlag(t *testing.T) {
	var (
		line = ` ping	127.0.0.1 -n = 4 --l =10 /i 5 `

		s      = []string{""}
		n      = []int{0, 0, 0}
		sCheck = []string{"127.0.0.1"}
		nCheck = []int{4, 5, 10}
	)

	argv := cmdline.SplitLine(line)
	flg := flag.NewFlagSet("cmdline", flag.PanicOnError)
	flg.StringVar(&s[0], "", "", "")
	flg.IntVar(&n[0], "n", -1, "count")
	flg.IntVar(&n[1], "i", -1, "TTL")
	flg.IntVar(&n[2], "l", -1, "size")
	flg.Parse(argv[1:])

	for i, v := range s {
		if v != sCheck[i] {
			t.Error(i, v, sCheck[i])
		}
	}
	for i, v := range n {
		if v != nCheck[i] {
			t.Error(i, v, nCheck[i])
		}
	}
}

//some complex case
func TestFlagFull(t *testing.T) {
	var (
		line = ` ping	 /l=2 127.0.0.1 --n	 1   ip2  -i=3 ip3 -r=	 5 -w =4 /k = 6 `

		s      = []string{"", "", ""}
		n      = []int{0, 0, 0, 0, 0, 0}
		sCheck = []string{"127.0.0.1", "ip2", "ip3"}
		nCheck = []int{1, 2, 3, 4, 5, 6}
	)

	argv := cmdline.SplitLine(line)
	flg := flag.NewFlagSet("cmdline", flag.PanicOnError)
	flg.StringVar(&s[0], "", "", "ip")
	flg.StringVar(&s[1], "", "", "ip2")
	flg.StringVar(&s[2], "", "", "ip3")
	flg.IntVar(&n[0], "n", -1, "count")
	flg.IntVar(&n[1], "l", -1, "size")
	flg.IntVar(&n[2], "i", -1, "TTL")
	flg.IntVar(&n[3], "w", -1, "timeout")
	flg.IntVar(&n[4], "r", -1, "count")
	flg.IntVar(&n[5], "k", -1, "host-list")
	flg.Parse(argv[1:])

	for i, v := range s {
		if v != sCheck[i] {
			t.Error(i, v, sCheck[i])
		}
	}
	for i, v := range n {
		if v != nCheck[i] {
			t.Error(i, v, nCheck[i])
		}
	}
}

//*/
