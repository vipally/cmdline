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
		n           = 0
	)
	cmdline.Summary("<thisCmd> is an example of command line usage to test gitbub.com/vipally.cmdline package.")
	cmdline.Details("no details")
	cmdline.CopyRight("no copyright in example cmd")

	cmdline.StringVar(&target_name, "", "target_name", "", true, "target host ip or name")
	cmdline.BoolVar(&v4, "4", "v4", v4, false, "ipv4")
	cmdline.IntVar(&ttl, "t", "ttl", ttl, false, "ttl")
	cmdline.IntVar(&ttl, "ttl", "synonym of -t", ttl, true, "this is synonym of -t, no mater thats this")

	cmdline.IntVar(&n, "c", "count", n, false, "count")
	cmdline.AnotherName("count", "c")
	cmdline.Parse()

	fmt.Println(cmdline.GetUsage())

	//Usage of ([ping] Build [Sep 29 2016 18:19:06]):
	//  Summary:
	//    <thisCmd> is an example of command line usage to test gitbub.com/vipally.cmdline package.

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
	//    no copyright in example cmd

	//  Details:
	//    no details
}
