package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"port-service/internal/config"
	"port-service/internal/repository/inmem"
	"port-service/internal/services"
	"port-service/internal/transport"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func run() error  {
	cfg := config.Read()

	portStoreRepo := inmem.NewPortStore()
	portService := services.NewPortService(portStoreRepo)

	// create http server with application injected
	httpServer := transport.NewHttpServer(portService)

	router := mux.NewRouter()
	router.HandleFunc("/ports", httpServer.GetPort).Methods("GET")
	router.HandleFunc("/count", httpServer.CountPorts).Methods("GET")
	router.HandleFunc("/ports", httpServer.UploadPorts).Methods("POST")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")

	return nil
}