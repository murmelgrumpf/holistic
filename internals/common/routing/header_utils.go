package routing

import "net/http"

func IsHtmx(header http.Header) bool {
	return header.Get("HX-Request") == "true"
}
