package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/shunsukenagashima/go_minting_api/src/gen/api"
	"github.com/shunsukenagashima/go_minting_api/src/handler"
	"golang.org/x/sync/errgroup"
)

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", 8888, err)
	}
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Printf("failed to load swagger %v", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	cn := &handler.CreateNft{}

	r := chi.NewRouter()

	// validation middlewareを使って、リクエストが
	// OpenAPIのスキーマに沿っているかチェックする
	r.Use(middleware.OapiRequestValidator(swagger))

	api.HandlerFromMux(cn, r)

	s := &http.Server{
		Handler: r,
		Addr: 	 fmt.Sprintf("0.0.0.0:%d", 8888),
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// context から終了通知を待機
	<- ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %+v", err)
	}

	eg.Wait()
}
