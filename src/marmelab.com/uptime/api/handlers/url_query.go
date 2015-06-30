package handlers

import (
	"strconv"
	"net/http"
	"log"
)
func parseQueryValues(req *http.Request, value string) int {
	val := req.URL.Query()[value]
	if val != nil {
		v, err := strconv.ParseInt(val[0], 10, 0)
		if err != nil {
			log.Print("error parseInt url query", err)
			return 0
		}
		return int(v)
	}
	return 0
}
