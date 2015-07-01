package handlers

import(
	"net/http"
)

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	SetCors(&header)
}
