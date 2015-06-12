package test

import (
	"../handlers"
	"net/http"
	"reflect"
	"testing"
)

func TestSetCorsShouldEffectivelySetCORSHeaders(t *testing.T) {
	expectedH := &http.Header{}
	expectedH.Set("Access-Control-Allow-Origin", "*")
	expectedH.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	expectedH.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	actualH := &http.Header{}
	handlers.SetCors(actualH)

	if !reflect.DeepEqual(expectedH, actualH) {
		t.Error("Error SetCors doesn't allow CORS")
	}
}
