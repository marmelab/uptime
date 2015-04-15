package main

import (
	"flag"
	"fmt"
	"time"
	"./poller"
)

func main() {
	request := poller.Request{}
	var duration int
	var err error
	dst := flag.String("dst", "8.8.8.8", "destination to ping")
	flag.Parse()
	fmt.Println("Ping on : " + *dst)
	request.SetdoPing(true)
	request.SettimeBetweenPing(1000000000)
	for request.DoPing {
		duration, err = poller.Ping(*dst)
		if duration != 0 || err == nil {
			fmt.Println("It works ! Time : ")
			fmt.Println(duration)
		} else {
			fmt.Println("It failed...")
			fmt.Println(err)
			break
		}
		time.Sleep(request.TimeBetweenPing)
	}

}