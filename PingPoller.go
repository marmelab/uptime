package main

import "golang.org/x/net/icmp"
import "os"
import "fmt"
import "net"


type Destination struct{
	url string
}



func Ping(s string) int{
	//var packetConn *icmp.PacketConn
	var errorCode int
	var data []byte
	ip,err := net.ResolveIPAddr("ip",s)
	if(err==nil){
		var dst net.Addr = ip
		packetConn, err := icmp.ListenPacket("ip4:icmp","")
		if(err==nil){
			errorCode,err := packetConn.WriteTo(data,dst)
			return errorCode
			if(err!=nil){
				fmt.Println("error on writeTo")
			}
		}else{
			fmt.Println("error on ListenPacket")
		}
	}else{
		errorCode := -1
		return errorCode
	}
	return errorCode
}

func main() {
	fmt.Println("Ping on : " + os.Args[1])
	errorCode :=Ping(os.Args[1])
	if(errorCode==0){
		fmt.Println("It works !")
	}else{
		fmt.Println("Ping failed....")
	}
}
