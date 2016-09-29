# cmdline [![GoDoc](https://godoc.org/github.com/vipally/cmdline?status.svg)](https://godoc.org/github.com/vipally/cmdline) ![Version](https://img.shields.io/badge/version-1.8.0-green.svg)
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
	8. Add synonyms support for flags
****

### usage of package cmdline
	
	cmdline.Version("1.0.2")
	cmdline.Summary("<thiscmd> is an example of cmdline package usage.")
	cmdline.Details(`Version   :<version>
    BulidTime :<buildtime>
    <thiscmd> is an example usage of github.com/vipally/cmdline package.`)
	cmdline.CopyRight("no copyright defined")

	//noname flag and require ones
	cmdline.StringVar(&target_name, "", "target_name", "", true, "target host ip or name")

	cmdline.BoolVar(&v4, "4", "v4", v4, false, "ipv4")

	//synonym with the same variables
	cmdline.IntVar(&ttl, "t", "ttl", ttl, false, "ttl")
	cmdline.IntVar(&ttl, "ttl", "synonym of -t", ttl, true, "this is synonym of -t")

	//define a synonym with method AnotherName
	c := cmdline.Int("c", "count", 0, false, "count")
	cmdline.AnotherName("count", "c")

	cmdline.Parse()

	fmt.Println(target_name, v4, ttl, *c)
	fmt.Println(cmdline.GetUsage())

	//output:
	//Usage of ([ping] Build [Sep 29 2016 21:14:37]):
	//  Summary:
	//    ping is an example of cmdline package usage.

	//  Usage:
	//    ping [-4=<v4>] [-c|count=<count>] [-t|ttl=<ttl>] <target_name>
	//  -4=<v4>	ipv4
	//  -c|count=<count>  int
	//      count
	//  -t|ttl=<ttl>  int
	//      ttl
	//  <target_name>  required  string
	//      target host ip or name

	//  CopyRight:
	//    no copyright defined

	//  Details:
	//    Version   :1.0.2
	//    BulidTime :[Sep 29 2016 21:14:37]
	//    ping is an example usage of github.com/vipally/cmdline package.
