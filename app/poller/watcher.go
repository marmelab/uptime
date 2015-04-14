package main

import (
"net/http" 
"fmt"	
"os"		
"strings"
)

type Request struct{
	url string
	method string
	headers []string
}

func main() {
	domainName := os.Getenv("NGINX_PORT_80_TCP")
	if(domainName != ""){
		domainName2 := strings.Split(domainName,"p")
		if(domainName2[1] != ""){
			resp,err:=sendPing("http"+domainName2[1])
			if(err==nil){
				fmt.Println("Status :" + resp.Status)
				fmt.Println("Protocole : " + resp.Proto)
			} else {
				fmt.Println(err)
			}
		}	
	}
}

func sendPing(domainName string) (resp *http.Response, err error){
	resp, err = http.Get(domainName)
	fmt.Println("Get sur " + domainName)
	return resp,err
}
