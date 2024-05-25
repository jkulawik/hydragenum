package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Server is starting")
	slog.Debug("mny benu")

	mux := http.NewServeMux()
	mux.HandleFunc("/test", testHandleFunc)

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
	}
}

func testHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("aaa")
}
