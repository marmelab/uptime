package poller

import (
	"golang.org/x/net/icmp"
	"net"
	"time"
	"io/ioutil"
	"encoding/json"
	"fmt"

)

type Response struct{
	Destination string
	Status string
	Time int
	Key string 
	Error error
}

func fromResponseToJSON(response Response) (data []byte, err error){
	return json.Marshal(response)
}

func RetrieveIpsFromJsonFile(fileName string) (data map[string]map[string]string){
	content,err := ioutil.ReadFile(fileName)
	if(err==nil){
		error := json.Unmarshal(content,&data)
		if(error!=nil){
			fmt.Println("error json")
		}
	}else{
		fmt.Println("error file")
	}
	return data
}

func FromDomainNameToIp(domainName string) (ip *net.IPAddr,err error){
	ip, err = net.ResolveIPAddr("ip", domainName) 
	return ip, err
}

func Ping(ip *net.IPAddr) (int, error) {
	var duration int
	var data []byte
		var destination net.Addr = ip
		packetConn, err := icmp.ListenPacket("ip4:icmp", "")
		if err == nil {
			timeNow := time.Now().Nanosecond()
			errorCode, err := packetConn.WriteTo(data, destination)
			duration = time.Now().Nanosecond() - timeNow
			if errorCode == 0 {
				return duration / 1000, err
			}
			if err != nil {
				return duration, err
			}
		} else {
			return duration, err
		}

	return duration / 1000, err
}
