package main

import (
	"net/http"

	"caddy-dash/ui"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /{$}", home)

	return mux
}
