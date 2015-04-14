package main

import (
"net/http/httptest"
 "net/http"
 "testing"
)

func TestsendPingFromWatcher200(t *testing.T) {
	resp := httptest.NewRecorder()
	mockServer:=httptest.newServer(http.HandlerFunc(func(w http.ReponseWriter, r *http.Request){
		}))
	defer mockServer.close()
	resp,err := watcher.sendPing(mockServer.URL)
	if(err==nil){
		fmt.Println("test")
	}else{
		t.Fail()
	}
}

func TestsendPingFromWatcher403(t *testing.T) {
	resp := httptest.NewRecorder()
	mockServer:=httptest.newServer(http.HandlerFunc(func(w http.ReponseWriter, r *http.Request){
		 w.WriteHeader().Set("Status-code","403")
		}))
	defer mockServer.close()
	resp,err := watcher.sendPing(mockServer.URL)
	if(err==nil){

	}else{
		t.Fail()
	}
}

func TestsendPingFromWatcher404(t *testing.T) {
	resp := httptest.NewRecorder()
	mockServer:=httptest.newServer(http.HandlerFunc(func(w http.ReponseWriter, r *http.Request){
		  w.WriteHeader().Set("Status-code","404")
		}))
	defer mockServer.close()
	resp,err := watcher.sendPing(mockServer.URL)
		if(err==nil){

	}else{
		t.Fail()
	}
}

func TestsendPingFromWatcher500(t *testing.T) {
	resp := httptest.NewRecorder()
	mockServer:=httptest.newServer(http.HandlerFunc(func(w http.ReponseWriter, r *http.Request){
		  w.WriteHeader().Set("Status-code","500")
		}))
	defer mockServer.close()
	resp,err := watcher.sendPing(mockServer.URL)
		if(err==nil){

	}else{
		t.Fail()
	}
}