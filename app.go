package main

import (
	"dark-web/server"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	_mux := server.New()
	_server := http.Server{
		Addr:    ":8080",
		Handler: _mux,
	}
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout, nil,
			),
		),
	)
	slog.Info(
		"Starting server...",
		"url",
		"http://localhost"+_server.Addr,
	)
	if err := _server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
