/*
 * @Author: Adrian Faisal
 * @Date: 31/08/21 1.00 PM
 */

package main

import (
	"context"
	"flag"
	"github.com/apldex/workshop-labti/internal/pkg/config"
	handlerProduct "github.com/apldex/workshop-labti/internal/pkg/handler/product"
	"github.com/apldex/workshop-labti/internal/pkg/resource/db"
	usecaseProduct "github.com/apldex/workshop-labti/internal/pkg/usecase/product"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.New(*configPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	persistentDB, err := db.NewPersistent(cfg.Database.Datasource)
	if err != nil {
		log.Fatalf("persistent db: %v", err)
	}

	productUsecase := usecaseProduct.NewUsecase(persistentDB)
	productHandler := handlerProduct.NewHandler(productUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK\n"))
	})

	r.HandleFunc("/product", productHandler.CreateProduct).Methods(http.MethodPost)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		osSignal := <-c

		log.Printf("got %s signal", osSignal.String())

		cancel()
	}()

	if err := startAPI(ctx, cfg.Server.Port, r); err != nil {
		log.Fatalf("start api failed: %v", err)
	}
}

func startAPI(ctx context.Context, addr string, handler http.Handler) error {
	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	log.Printf("server running at %s", addr)

	// wait for context cancellation
	<-ctx.Done()

	log.Printf("shutting down server...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("shutdown server failed: %v", err)
	}

	log.Printf("server stopped gracefully.")
	return nil
}
