package main

import (
	"fmt"

	"github.com/vipally/cmdline"
)

func main() {
	var (
		target_name string
		v4          = false
		ttl         = 0
	)
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
}
