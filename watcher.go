package main

import "net/http"
import "fmt"

func main() {
	sendPing()
}

func sendPing() {
	resp, err := http.Get("http://google.fr")
	fmt.Println("Get on google.fr :")
	if(err==nil){
		fmt.Println("Status :" + resp.Status)
		fmt.Println("Protocole : " + resp.Proto)
	}
}