package poller

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type Response struct {
	Id int
	Destination string
	Status      string
	Time        int
	Created_at  time.Time
	Target_id   int
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

func RetrieveConfDbFromJsonFile(fileName string) (data map[string]interface{}) {
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

func HttpPing(url string, protocol string) (int, error) {
	if url == "" {
		error := errors.New("url is null")
		return -1, error
	}

	log.Printf("Pinging %v", url)

	var duration int
	timeNow := time.Now().Nanosecond()
	_, err := http.Get(protocol + "://" + url)
	duration = time.Now().Nanosecond() - timeNow
	duration = duration / 1000
	if err != nil {
		log.Printf("Error %v", err)
		return duration, err
	}

	log.Printf("Duration: %vms", duration)

	return duration, err
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
			log.Print("icmp ListenPacket error ")
		}
	}
	errorCode, err := packetConn.WriteTo(data, ip)
	duration = time.Now().Nanosecond() - timeNow
	if errorCode == 0 {
		return duration / 1000, nil
	}
	if err != nil {
		return duration, err
	}
	return duration / 1000, err
}
