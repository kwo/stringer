package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/kwo/stringer/api/greader"
	"github.com/kwo/stringer/repository/bogus"
)

func main() {
	var mainErr error
	mainCtx, mainCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer mainCancel()

	// internal service configuration
	bogusProvider := bogus.New()
	greaderHandler := greader.New(bogusProvider)

	// httpd configuration
	handler := chi.NewMux()
	handler.Mount("/reader", greaderHandler)

	httpd := &http.Server{
		Addr:    "localhost:8888",
		Handler: handler,
	}
	go func() {
		if err := httpd.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("cannot start httpd: %s", err)
			mainErr = err
			mainCancel()
		}
	}()
	log.Printf("http listening on %s ...", httpd.Addr)

	// initialization complete, wait for shutdown signal

	<-mainCtx.Done()
	log.Println("stopping...")

	// stop httpd
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := httpd.Shutdown(ctx); err != nil {
		log.Printf("http shutdown error: %s", err)
	}

	// done
	if mainErr != nil {
		log.Fatal("done1")
	}
	log.Println("done")
}
