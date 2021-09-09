package main

import (
	"net/http"
	"os"
	"wing/initialize"
)

const (
	defaultPort = "9200"
)

func getPort() (serverPort string) {
	port := os.Getenv("SERVERPORT")
	if port == "" {
		return defaultPort
	}

	return port
}

func main() {

	srv := &http.Server{
		Addr:    "0.0.0.0:" + getPort(),
		Handler: initialize.Initialize(),
	}

	srv.ListenAndServe()
}
