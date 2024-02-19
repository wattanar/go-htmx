package main

import (
	"context"
	"fmt"
	"go-htmx/internal/pkg/logger"
	"go-htmx/internal/pkg/recover"
	"go-htmx/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer recover.Recover()

	router := setupRouter()
	logger := logger.NewLogger()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", "8080"),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		logger.Info("Starting server...", "port", server.Addr)
		fmt.Println(server.ListenAndServe())
	}()

	// Gracefully Shutdown
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server %s", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}

func setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	routes.NewRouter(router)
	return router
}
