package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/Faroukhamadi/farouk-blog-api/api"
)

func main() {
	s := api.Init()

	go func() {
		log.Println("Starting server on port", s.Addr)
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)

	sig := <-q
	log.Println("Got signal", sig, "to shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
