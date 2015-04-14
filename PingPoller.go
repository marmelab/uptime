package main

import "golang.org/x/net/icmp"
import "net"
import "os"
import "fmt"


type Destination struct{
	url string
}

func () String() {
	
}

func Ping(s string) int{
	var packetConn *icmp.PacketConn
	var data []byte
	var dst net.Addr = s
	packetConn = icmp.ListenPacket("ip4:icmp","")
	duration,err = packetConn.WriteTo(data,dst)
	return duration
}

func main() {
	duration :=Ping(os.Args[1])
	fmt.Println(duration)
}
