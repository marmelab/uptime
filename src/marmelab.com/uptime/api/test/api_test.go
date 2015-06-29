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

func TestSetContentTypeShouldEffectivelySetHeader(t *testing.T) {
	expectedH := &http.Header{}
	expectedH.Set("Content-Type", "application/json")

	actualH := &http.Header{}
	handlers.SetContentType(actualH)

	if !reflect.DeepEqual(expectedH, actualH) {
		t.Error("Error SetContentType doesn't set Content-Type as json")
	}
}

func TestSetContentTypeWorksForGetRequest() {

}
