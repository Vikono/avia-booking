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
	"golang.org/x/sync/errgroup"

	v1 "homework/internal/api/v1"
	"homework/internal/config"
	"homework/internal/db"
	"homework/specs"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // Call the next handler
		log.Printf("Request: %s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func main() {
	var (
		err         error
		ctx, cancel = signal.NotifyContext(
			context.Background(),
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
	)
	defer cancel()

	cfg, err := config.InitConfig(os.Args)
	if err != nil {
		log.Fatal("get config: ", err.Error())
		return
	}

	DB := db.Connect()
	fmt.Println("???/")
	// defer db.CloseConnection(DB)

	apiServer := v1.NewAPIServer(DB)
	fmt.Println("???/")

	err = startHTTPServer(ctx, cfg, apiServer)
	fmt.Println("Statring server...")
	if err != nil {
		log.Fatal("starting server: ", err.Error())
	}
}

func startHTTPServer(
	ctx context.Context,
	cfg *config.Config,
	apiServer specs.ServerInterface,
	middlewares ...specs.MiddlewareFunc,
) error {
	handler := specs.HandlerWithOptions(apiServer, specs.ChiServerOptions{
		BaseURL:     cfg.BasePath,
		Middlewares: middlewares,
	})

	router := chi.NewRouter()
	router.Handle("/*", loggingMiddleware(handler))

	httpServer := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return httpServer.Shutdown(ctx)
	})

	return group.Wait()
}
