# cmdline [![GoDoc](https://godoc.org/github.com/vipally/cmdline?status.svg)](https://godoc.org/github.com/vipally/cmdline) ![Version](https://img.shields.io/badge/version-1.2.0-green.svg)
	cmdline is a Golang package based on std.flag.
	It extend the std flag package and improve the user interface and add nessary usage message.
****

	Ally(vipally@gmail.com) modify from std.flag version 1.7
	1. Add LogicName and Required field for every flag, and modify the flag define interface
	2. Add Summary and Details for command line info
	3. Add interface GetUsage() string
	4. Modify the Parse() logic
	5. Add noname-flag support
	6. Add /flag support
	7. Fix "-flag = x" or "-flag= x" or "-flag =x" cause panic bug
****

	//usage of cmdline as follow
	func main() {
		cmdline.Summary("command <thiscmd> is used to copy a file to another path.")
		cmdline.Details(`Command <thiscmd> is used to copy a file to another path.
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
