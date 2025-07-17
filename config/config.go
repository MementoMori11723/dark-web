package config

import (
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

//go:embed tor/vanity/hostname
var url string

func New() string {
	_port := os.Getenv("PORT")
	if _port == "" {
		_port = "80"
	}
	_port = fmt.Sprintf(
		":%s", _port,
	)
	slog.SetDefault(
		slog.New(
			slog.NewJSONHandler(
				os.Stdout, nil,
			),
		),
	)
	return _port
}

func GetUrl() string {
	return fmt.Sprintf(
		"http://%s", strings.ReplaceAll(
			url, "\n", "",
		),
	)
}
