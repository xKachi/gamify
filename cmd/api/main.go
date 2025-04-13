package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// API version
const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {

	// config instance
	var cfg config

	// commandline flags
	flag.IntVar(&cfg.port, "port", 4000, "server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// application struct instance
	app := &application{
		config: cfg,
		logger: logger,
	}

	// routes handler
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//server
	srv := &http.Server{
		// "Format this as a string that starts with a colon followed by a decimal number.Output: :8080"
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the server
	logger.Info("Starting Server", "addr", srv.Addr, "env", cfg.env)
	err := srv.ListenAndServe()
	// Convert error to human readable format
	logger.Error(err.Error())
	// terminates program
	os.Exit(1)

}
