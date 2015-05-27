package test

import (
	"../../api"
	"net/http"
	"reflect"
	"testing"
)

func TestSetCorsShouldNotTriggerError(t *testing.T) {
	exepectedH := &http.Header{}
	exepectedH.Set("Access-Control-Allow-Origin", "*")
	exepectedH.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	exepectedH.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	actualH := &http.Header{}
	main.SetCors(actualH)

	if !reflect.DeepEqual(exepectedH, actualH) {
		t.Error("Error SetCors don't allow Cors")
	}
}
