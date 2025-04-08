package main

import (
	"holistic/api"
	"holistic/internals/common/logger"
	"holistic/internals/common/logging"
	"holistic/internals/routing"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type Router struct {
	*chi.Mux
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	logging.Logo()
	router := chi.NewMux()

	router.Handle("/*", routing.HandleNotFound())
	logger.Info("404 Redirect registered")

	router.Handle("/public/*", public())
	logger.Info("Public Endpoints registered")

	api.InitializeEndpoints(router)
	logger.Info("Router Endpoints registered")

	listenAddr := os.Getenv("LISTEN_ADDR")
	logger.Info("HTTP server started", "listenAddr", listenAddr)

	logging.Delimiter()
	http.ListenAndServe(listenAddr, router)
}
