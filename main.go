package main

import (
	"net/http"
	"time"
)

func main() {
	// She's alive
	p("Flare", version(), "started at", config.Address)

	// Create a multiplexer for handling requests
	mux := http.NewServeMux()

	// Handle static files
	staticFileServer := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	// Main page
	mux.HandleFunc("/", index)
	// Error page
	mux.HandleFunc("/err", err)

	// Configuring and starting the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

	return
}
