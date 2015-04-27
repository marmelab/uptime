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
<<<<<<< HEAD
<<<<<<< HEAD
func Ping(ip *net.IPAddr, packetConn *icmp.PacketConn) (int, error) {
=======
=======
>>>>>>>  code and test fixed
func Ping(ip *net.IPAddr) (int, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> test failed again
=======
>>>>>>> tests fixed
=======
func Ping(ip *net.IPAddr) (int, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> test failed again
=======
>>>>>>> tests fixed
=======
<<<<<<< HEAD
func Ping(ip *net.IPAddr) (int, error) {
<<<<<<< HEAD
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	 if ip == nil && &ip != nil {
	 	error := errors.New("ip = nil ")
	 	return 0, error
	 }
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>> test failed again
=======
>>>>>>> test failed again
=======
>>>>>>> tests fixed
=======
=======
>>>>>>> tests fixed
>>>>>>> tests fixed
=======
=======
=======
>>>>>>> tests fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
>>>>>>> test failed again
=======
=======
>>>>>>> tests fixed
<<<<<<< HEAD
=======
=======
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
>>>>>>> tests fixed
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> test failed again
=======
>>>>>>>  code and test fixed
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
func Ping(ip *net.IPAddr, packetConn *icmp.PacketConn) (int, error) {
=======
=======
func Ping(ip *net.IPAddr, packetConn *icmp.PacketConn) (int, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  code and test fixed
=======
>>>>>>>  add instance of packetConn if it is nil
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	if packetConn == nil {
		return -1, errors.New("error argument packetConn nil")
	}
>>>>>>>  code and test fixed
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  add instance of packetConn if it is nil
=======
>>>>>>>  code and test fixed
=======
=======
>>>>>>>  add instance of packetConn if it is nil
>>>>>>>  add instance of packetConn if it is nil
=======
=======
>>>>>>>  add instance of packetConn if it is nil
=======
=======
=======
>>>>>>> d09aa0ab1d64898cc222c40cd23103b41a95c1e5
func Ping(ip *net.IPAddr, packetConn *icmp.PacketConn) (int, error) {
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	if ip == nil {
		return -1, errors.New("error argument ip nil")
	}
	var duration int
	var data []byte
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>>  add instance of packetConn if it is nil
=======
>>>>>>>  code and test fixed
=======
>>>>>>>  add instance of packetConn if it is nil
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	var err error
	timeNow := time.Now().Nanosecond()
	if packetConn == nil {
		packetConn,err = icmp.ListenPacket("ip4:icmp", "")
		if(err != nil){
			return -1,err
		}
	}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
	timeNow := time.Now().Nanosecond()
>>>>>>>  code and test fixed
=======
<<<<<<< HEAD
>>>>>>>  add instance of packetConn if it is nil
=======
=======
	timeNow := time.Now().Nanosecond()
>>>>>>>  code and test fixed
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
=======
=======
=======
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
	var err error
	timeNow := time.Now().Nanosecond()
	if packetConn == nil {
		packetConn,err = icmp.ListenPacket("ip4:icmp", "")
		if(err != nil){
			return -1,err
		}
	}
<<<<<<< HEAD
>>>>>>>  add instance of packetConn if it is nil
=======
<<<<<<< HEAD
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
>>>>>>>  add instance of packetConn if it is nil
	_, _ = packetConn.WriteTo(data, ip)
	duration = time.Now().Nanosecond() - timeNow
	return duration / 1000, nil
>>>>>>>  code and test fixed
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
>>>>>>>  add instance of packetConn if it is nil
=======
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
=======
>>>>>>> test failed again
=======
>>>>>>>  code and test fixed
=======
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
=======
>>>>>>> test failed again
>>>>>>> test failed again
=======
>>>>>>>  code and test fixed
=======
	_, _ = packetConn.WriteTo(data, ip)
	duration = time.Now().Nanosecond() - timeNow
	return duration / 1000, nil
<<<<<<< HEAD
>>>>>>>  code and test fixed
=======
>>>>>>> d09aa0ab1d64898cc222c40cd23103b41a95c1e5
>>>>>>> fe35ab3749b7451f789da23b18b4944146380c19
>>>>>>> 5e669e907a348765a72d9c371814de2c87ae53df
}
