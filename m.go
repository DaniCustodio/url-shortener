package main

import (
	"log/slog"
	"main/api"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		return
	}
	slog.Info("run completed")
}

// The main function can't return an error, so we need to call a function that can
// This facilitates logging and error handling
func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
