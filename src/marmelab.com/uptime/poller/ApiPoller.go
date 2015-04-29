package poller

import (
	"bytes"
	"net/http"
	"encoding/json"
)

func DoPostOn(response *Response, url string) error {
	 data,_ := json.Marshal(response)
	req, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		return error
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
