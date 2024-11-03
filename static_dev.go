//+build dev
//go:build dev
// +build dev

package main

import (
	"holistic/internals/common/logger"

	"net/http"
	"os"
)

func public() http.HandlerFunc {
	logger.Info("Building static files for development")
	fs := http.StripPrefix("/public/", http.FileServerFS(os.DirFS("api/public")))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("max-age", "0")
		fs.ServeHTTP(w, r)
	}
}
