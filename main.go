package main

import (
	"holistic/api/handlers"
	"holistic/internals/common/logger"
	"log"
	"log/slog"
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

	logger.Logo()

	router := chi.NewMux()
	router.Handle("/*", public())
	slog.Info("Public Endpoints registered")

	handlers.Handle(router)
	slog.Info("Router Endpoints registered")

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)

	logger.Delimiter()
	http.ListenAndServe(listenAddr, router)
}
