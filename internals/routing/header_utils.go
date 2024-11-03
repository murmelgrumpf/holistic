package routing

import "net/http"

func IsHtmx(header http.Header) bool {
	return header.Get("HX-Request") == "true"
}

func IsDocument(header http.Header) bool {
	return header.Get("Sec-Fetch-Dest") == "document"
}
