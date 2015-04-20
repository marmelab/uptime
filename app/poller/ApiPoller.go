package poller

import (
	"bytes"
	"net/http"
	"strconv"
)

func DoPostOn(response *Response, url string)error{
	var result = []byte(`{"Destination": "` + response.Destination + `","Time": ` + strconv.Itoa(response.Time) + `,"Status": "` + response.Status + `"}`)
	req, error := http.NewRequest("POST", url, bytes.NewBuffer(result))
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
