package cmdline_test

import (
	"github.com/vipally/cmdline"
)

func ExampleCmdline() {
	cmdline.Summary("command copy is used to copy a file to another path.")
	cmdline.Details(`Command copy is used to copy a file to another path.
    If the destnation file is exist, default ask for if will cover it.
    If flag -y used, it will cover the destnation file without ask.
    If flag -n used, it will not cover the destnation file without ask.
`)
	cmdline.String("s", "src", ".", true, "source file path")
	cmdline.String("d", "dst", ".", true, "destnation file path")
	cmdline.Bool("c", "cover", false, false, "if cover the destnation file")
	cmdline.Bool("y", "yes", false, false, "if auto select yes when ask for cover")
	cmdline.Bool("n", "no", false, false, "if auto select no when ask for cover")
	cmdline.Parse()

	//[error] require but lack of flag -s=<src>
	//Usage of [copy.exe]:
	//  Summary:
	//    command copy is used to copy a file to another path
	//
	//  Usage:
	//    copy.exe [-c=<cover>] -d=<dst> [-n=<no>] -s=<src> [-y=<yes>]
	//  -c=<cover>
	//      if cover the destnation file
	//  -d=<dst>  required  string (default ".")
	//      destnation file path
	//  -n=<no>	if auto select no when ask for cover
	//  -s=<src>  required  string (default ".")
	//      source file path
	//  -y=<yes>
	//      if auto select yes when ask for cover
	//
	//  Details:
	//    Command copy is used to copy a file to another path.
	//    If the destnation file is exist, default ask for if will cover it.
	//    If flag -y used, it will cover the destnation file without ask.
	//    If flag -n used, it will not cover the destnation file without ask.
}
