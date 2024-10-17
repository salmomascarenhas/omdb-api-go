package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/salmomascarenhas/omdb-api-go/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute application: ", "error", err)
		return
	}
	slog.Info("All done")
}

func run() error {
	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OMDB_API_KEY is not set")
	}

	handler := api.NewHandler(apiKey)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":3000",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
