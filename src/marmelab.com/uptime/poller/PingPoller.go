package poller

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"io/ioutil"
	"net"
	"time"
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

func Ping(ip *net.IPAddr, packetConn *icmp.PacketConn) (int, error) {
	if ip == nil && &ip != nil {
		error := errors.New("ip = nil ")
		return 0, error
	}
	var duration int
	var data []byte
	var err error
	timeNow := time.Now().Nanosecond()
	if packetConn == nil {
		packetConn, err = icmp.ListenPacket("ip4:icmp", "")
		if err != nil {
			log.Print("erreur dans nv packet")
			return -1, err
		}
	}
	errorCode, err := packetConn.WriteTo(data, ip)
	duration = time.Now().Nanosecond() - timeNow
	if errorCode == 0 {
		log.Print("erreur dans errcode")
		return duration / 1000, nil
	}
	if err != nil {
		log.Print("erreur dans err ping")
		return duration, err
	}
	log.Print("pas d'erreur")
	return duration / 1000, err
}
