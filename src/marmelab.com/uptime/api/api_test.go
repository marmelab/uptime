package poller

import (
	"golang.org/x/net/icmp"
	"net"
	"testing"
)

type myHttp interface {
	ResponseWriter
}

type myHttp.ResponseWriter interface {

}
func TestsetAllowCORSwithValidResponseWriterShouldNotTriggerError(t *testing.T) {
	w = new myHttp.ResponseWriter{}
	err := setAllowCORS(&w)
	if err != nil {
		t.Error("allow CORS on a valid ResponseWriter should not raise a error, got ", err)
	}
}

func TestsetAllowCORSwithInvalidResponseWriterShouldTriggerError(t *testing.T) {
	err := setAllowCORS(nil)
	if err == nil {
		t.Error("allow CORS on a invalid ResponseWriter should raise a error, got ", err)
	}
}

func TestsetAllowCORSwithValidResponseWriterShouldHaveCORSAllowed(t *testing.T) {
	err := 
	if err != nil {
		t.Error("ResponseWriter should have CORS allowed on his header, got ", err)
	}
}

func Testreturn500WithValidResponseWriterShouldNotTriggerError(t *testing.T) {
	w = new myHttp.ResponseWriter{}
	err := return500(&w)
	if err != nil {
		t.Error("using return500 with a valid ResponseWriter should not raise a error, got ", err)
	}
}

func Testreturn500WithInValidResponseWriterShouldTriggerError(t *testing.T) {
	err := return500(nil)
	if err == nil {
		t.Error("using return500 with a valid ResponseWriter should raise a error, got ", err)
	}
}

func Testreturn500WithValidResponseWriterShouldReturn500(t *testing.T) {
	_, err := 
	if err != nil {
		t.Error("ResponseWriter should have 500 as status code , got ", err)
	}
}


