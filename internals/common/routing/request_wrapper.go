package routing

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

func Make(endpoint Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestContext context.Context
		switch endpoint.Request {
		case POST, PUT, PATCH:
			requestContext = r.Context()
		}

		response, err := endpoint.Function(r.Header, requestContext)
		if err != nil {
			slog.Error("HTTP API error", "err", err, "path", r.URL.Path)
			if apiErr, ok := err.(APIError); ok {
				if err := writeJson(w, apiErr.StatusCode, apiErr); err != nil {
					internalError(w)
					slog.Error("HTTP API error", "err", err, "path", r.URL.Path)
				}
			} else {
				internalError(w)
			}
			return
		}

		switch endpoint.Response {
		case JSON:
			if err := writeJson(w, http.StatusOK, response); err != nil {
				internalError(w)
				slog.Error("Converting response to JSON error", "err", err, "path", r.URL.Path)
			}

		case TEMPL:
			if template, ok := response.(templ.Component); ok {
				if err := template.Render(r.Context(), w); err != nil {
					internalError(w)
					slog.Error("Rendering Templ Component error", "err", err, "path", r.URL.Path)
				}
			} else {
				internalError(w)
				slog.Error("Converting response to Templ Component error", "err", err, "path", r.URL.Path)
			}

		}

	}
}

func Redirect(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, http.StatusSeeOther)
	}
}

func internalError(w http.ResponseWriter) {
	errResp := map[string]any{
		"statusCode": http.StatusInternalServerError,
		"msg":        "internal server error",
	}
	writeJson(w, http.StatusInternalServerError, errResp)
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}
