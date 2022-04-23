package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"

	"github.com/kwo/stringer/api/greader"
	"github.com/kwo/stringer/repository/bogus"
)

func main() {

	bogusProvider := bogus.New()
	greaderHandler := greader.New(bogusProvider)

	handler := chi.NewMux()
	handler.Mount("/greader", greaderHandler)

	httpd := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: handler,
	}

	ctx, killSwitch := context.WithCancel(context.Background())
	defer killSwitch()

	go func() {
		if err := httpd.ListenAndServe(); err != nil {
			log.Printf("cannot start httpd: %s", err)
			killSwitch()
		}
	}()
	log.Println("started...")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
		fmt.Println()
		log.Println("caught signal")
	case <-ctx.Done():
		log.Println("context cancelled")

	}
	log.Println("stopping...")

	// TODO: wrap in function
	dbCtx, dbCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer dbCancel()
	if err := httpd.Shutdown(dbCtx); err != nil {
		log.Printf("error shutting down server: %s", err)
		if errClose := httpd.Close(); errClose != nil {
			log.Printf("server close failed: %s", errClose)
		}
	}

	log.Println("exit")

}
