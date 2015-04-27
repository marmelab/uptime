package poller

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/icmp"
	"io/ioutil"
	"net"
	"time"
	"errors"
	"log"
)

type Response struct {
	Destination string
	Status      string
	Time        int
	Key         string
	Error       error
}

func RetrieveIpsFromJsonFile(fileName string) (data map[string]string) {
	content, err := ioutil.ReadFile(fileName)
	if err == nil {
		error := json.Unmarshal(content, &data)
		if error != nil {
			fmt.Println("error Unmarshal json : ", error)
		}
	} else {
		fmt.Println("error read file : ", err)
	}
	return data
}

func FromDomainNameToIp(domainName string) (ip *net.IPAddr, err error) {
	return net.ResolveIPAddr("ip", domainName)
}

func Ping(ip *net.IPAddr) (int, error) {
	if ip == nil {
		log.Print(&ip)
		error := errors.New("ip = nil ")
		return 0, error
	}
	var duration int
	var data []byte
	packetConn, err := icmp.ListenPacket("ip4:icmp", "")
	if err == nil {
		timeNow := time.Now().Nanosecond()
		errorCode, err := packetConn.WriteTo(data, ip)
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
