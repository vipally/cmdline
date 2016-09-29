//this example shows main features of package cmdline.
package main

import (
	"fmt"

	"github.com/vipally/cmdline"
)

func main() {
	var (
		host string
		v4   = false
		ttl  = 0
	)
	cmdline.Version("1.0.2")
	cmdline.Summary("<thiscmd> is an example of cmdline package usage.")
	cmdline.Details(`Version   :<version>
    BulidTime :<buildtime>
    <thiscmd> is an example usage of github.com/vipally/cmdline package.`)
	cmdline.CopyRight("no copyright defined")

	//no-name flag and required ones
	cmdline.StringVar(&host, "", "host", "", true, "host ip or name")
	host2 := cmdline.String("", "host2", "", false, "second host ip or name")

	cmdline.BoolVar(&v4, "4", "v4", v4, false, "ipv4")

	//synonym with the same variables
	cmdline.IntVar(&ttl, "t", "ttl", ttl, false, "ttl")
	cmdline.IntVar(&ttl, "ttl", "synonym of -t", ttl, true, "this is synonym of -t")

	//define a synonym with method AnotherName
	c := cmdline.Int("c", "count", 0, false, "count")
	cmdline.AnotherName("count", "c")

	cmdline.Parse()

	fmt.Printf("host:%s host2:%s v4:%t ttl:%d count:%d\n", host, *host2, v4, ttl, *c)
	fmt.Println(cmdline.GetUsage())

	//cmd example: ping -t=20 /4 127.0.0.1 --count =4 localhost -ttl= 5
	//output:
	//host:127.0.0.1 host2:localhost v4:true ttl:5 count:4
	//Usage of ([ping] Build [Sep 29 2016 23:50:04]):
	//  Summary:
	//    ping is an example of cmdline package usage.
	//
	//  Usage:
	//    ping [-4=<v4>] [-c|count=<count>] [-t|ttl=<ttl>] <host> [<host2>]
	//  -4=<v4>       ipv4
	//  -c|count=<count>  int
	//      count
	//  -t|ttl=<ttl>  int
	//      ttl
	//  <host>  required  string      host ip or name
	//  <host2>  required  string     second host ip or name
	//
	//  CopyRight:
	//    no copyright defined
	//
	//  Details:
	//    Version   :1.0.2
	//    BulidTime :[Sep 29 2016 23:50:04]
	//    ping is an example usage of github.com/vipally/cmdline package.
}
