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
	"github.com/Rishipreeth03/Api_Handler/Api_Handler/internal/http/handlers/student"
	"github.com/Rishipreeth03/Api_Handler/Api_Handler/internal/storage/sqlite"
)

func main() {
	//loadconfig
	cfg := config.MUSTLoad()
	//database setup
	//you can simply shift to other storage by changing this line
	//like	postgres.New(cfg) or mongo.New(cfg) and do
	//not forget to implement storage.Storage interface methods
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Storage Initialised successfully", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))

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
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown", slog.String("error", err.Error()))
	}
	slog.Info("Server exited properly")
}
