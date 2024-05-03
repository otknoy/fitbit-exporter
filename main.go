package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	slog.Info("start fitbit-exporter")
	defer slog.Info("end fitbit-exporter")

	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", nil); err != http.ErrAbortHandler {
		slog.Error("error: %s", err)
	}

	fmt.Println("hello")
}
