package poller

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"io/ioutil"
	"net"
	"time"
)

type Response struct {
	Destination string
	Status      string
	Time        int
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

<<<<<<< HEAD
func Ping(ip *net.IPAddr) (int, error) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> test failed again
	 if ip == nil && &ip != nil {
	 	error := errors.New("ip = nil ")
	 	return 0, error
	 }
<<<<<<< HEAD
=======
>>>>>>> tests fixed
=======
>>>>>>> test failed again
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
<<<<<<< HEAD
=======
func Ping(ip *net.IPAddr, packetConn *icmp.PacketConn) (int, error) {
	if ip == nil {
		return -1, errors.New("error argument ip nil")
	}
	var duration int
	var data []byte
	var err error
	timeNow := time.Now().Nanosecond()
	if packetConn == nil {
		packetConn,err = icmp.ListenPacket("ip4:icmp", "")
		if(err != nil){
			return -1,err
		}
	}
	_, _ = packetConn.WriteTo(data, ip)
	duration = time.Now().Nanosecond() - timeNow
	return duration / 1000, nil
>>>>>>>  code and test fixed
=======
>>>>>>> test failed again
}
