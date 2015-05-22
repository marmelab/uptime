package test

import (
	"testing"
	"../../api"
)

// w *http.ResponseWriter, ne pas passer cet objet mais une interface ?

type Header map[string][]string

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int,error)
	WriteHeader(int)
}

func TestSetAllowCORSwithValidResponseWriterShouldNotTriggerError(t *testing.T) {
	w :=  *ResponseWriter{}
	err := main.SetAllowCORS(&w)
	if err != nil {
		t.Error("allow CORS on a valid ResponseWriter should not raise a error, got ", err)
	}
}

func TestSetAllowCORSwithInvalidResponseWriterShouldTriggerError(t *testing.T) {
	err := main.SetAllowCORS(nil)
	if err == nil {
		t.Error("allow CORS on a invalid ResponseWriter should raise a error, got ", err)
	}
}

/*func TestSetAllowCORSwithValidResponseWriterShouldHaveCORSAllowed(t *testing.T) {
	err := 
	if err != nil {
		t.Error("ResponseWriter should have CORS allowed on his header, got ", err)
	}
}*/

func TestReturn500WithValidResponseWriterShouldNotTriggerError(t *testing.T) {
	w := ResponseWriter{}
	err := main.Return500(&w)
	if err != nil {
		t.Error("using return500 with a valid ResponseWriter should not raise a error, got ", err)
	}
}

func TestReturn500WithInValidResponseWriterShouldTriggerError(t *testing.T) {
	err := main.Return500(nil)
	if err == nil {
		t.Error("using return500 with a valid ResponseWriter should raise a error, got ", err)
	}
}

/*func TestReturn500WithValidResponseWriterShouldReturn500(t *testing.T) {
	_, err := 
	if err != nil {
		t.Error("ResponseWriter should have 500 as status code , got ", err)
	}
}*/


