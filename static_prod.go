//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed api/public
var publicFS embed.FS

func public() http.HandlerFunc {
	fs := http.StripPrefix("/public/", http.FileServerFS(publicFS))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("max-age", "0")
		fs.ServeHTTP(w, r)
	}
}
