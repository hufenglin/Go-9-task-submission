package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	serverOut := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8086",
	}

	group.Go(func() error {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Start server error,will exit.", err.Error())
		}
		return err
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("Errgroup exit...", ctx.Err().Error())
		case <-serverOut:
			fmt.Println("Http Request shutdown, server will exit...")
		}

		fmt.Println("Shutting down server...")
		err := server.Shutdown(ctx)
		return err
	})

	group.Go(func() error {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			fmt.Println("Errgroup exit,", ctx.Err().Error())
			return ctx.Err()
		case sig := <-sigChan:
			fmt.Println("Caught exit reason ", sig.String())
			return fmt.Errorf("%s\n", sig.String())
		}
	})

	fmt.Printf("Finished, errgroup exiting, %+v\n", group.Wait())
}
