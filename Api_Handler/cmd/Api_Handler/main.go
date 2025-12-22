package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rishipreeth03/Api_Handler/Api_Handler/internal/config"
)

func main() {
	//loadconfig
	cfg := config.MUSTLoad()
	//database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to API Handler"))
	})
	//setup server

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}
	slog.Info("Server started %s", slog.String("address", cfg.HTTPServer.Addr))
	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	slog.Info("Server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown", slog.String("error", err.Error()))
	}
	slog.Info("Server exited properly")
}
