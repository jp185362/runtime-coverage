package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"os/signal"
	"e2e/msg"
)
const address = ":8080"
const timedShutdown = true
const lifetimeSeconds = 10

func main() {
	log.Println("Initializing http Server on address", address)

	mux := http.NewServeMux()
	mux.Handle("/conditional", msg.ConditionalHandler())
	mux.Handle("/basic", msg.MsgHandler())
	server := &http.Server{Addr: address, Handler: mux}

	// Run the server in a background Go routine
    go func() {
        if err := server.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

	// Wait for shutdown
	waitShutdown()

	log.Println("Shutting down server...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
	log.Fatal("Failed to shut down server safely", server.Shutdown(ctx))
}

func waitShutdown() {
	if timedShutdown {
		log.Println("Waiting for service lifetime of", lifetimeSeconds, "seconds")
		<-time.After(lifetimeSeconds * time.Second)
	} else {
		// Setting up signal capturing
		// Waiting for SIGINT (kill -2) or timeout
		log.Println("Waiting for shutdown signal")
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop
	}
}