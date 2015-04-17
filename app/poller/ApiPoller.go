package poller

import (
    "log"
    "net/http"
    "bytes"
)

func DoPostOn(response *Response,url string,) {
	var result = []byte(`{"Destination": "`+response.Destination+`"}`)
	req,error :=http.NewRequest("POST",url,bytes.NewBuffer(result))
	if(error!=nil){
		log.Fatal(error)
	}
	client := &http.Client{}
	resp,err := client.Do(req)
	if(err!=nil){
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
