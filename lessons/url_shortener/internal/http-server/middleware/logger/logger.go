package logger

import (
	"log/slog"
	"net/http"
)

func New(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log = log.With()
	}

}
