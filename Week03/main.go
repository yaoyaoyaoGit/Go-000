package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

var sigs chan os.Signal

type server struct {
	*http.Server

	name string
}

type handler struct {
	name string
}

func main() {
	sigs = make(chan os.Signal, 1)
	group, ctx := errgroup.WithContext(context.Background())

	s1 := newServer("1", ":8080")
	s2 := newServer("2", ":9090")

	group.Go(
		func() error {
			// If server is closed by close or shutdown, we don't return error so that err in waitGroup won't be polluted.
			if err := s1.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}
			return nil
		},
	)
	group.Go(
		func() error {
			if err := s2.ListenAndServe(); err != http.ErrServerClosed {
				return err
			}
			return nil
		},
	)

	group.Go(func() error {
		return listenSignal(ctx, s1, s2)
	})

	if err := group.Wait(); err != nil {
		fmt.Println("1", err)
	}

}

func (h handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Server %s received request", h.name)
}

func newServer(name, addr string) server {
	h := handler{name}

	return server{
		&http.Server{
			Addr:    addr,
			Handler: h,
		},
		name,
	}
}

func listenSignal(ctx context.Context, s1, s2 server) error {
	signal.Notify(sigs)
	select {
	case sig := <-sigs:
		fmt.Printf("Recevied signal %s\n", sig.String())
		// use context.Backgroud to shutdown the server because original ctx is canceled
		if err := s1.Shutdown(context.Background()); err != nil {
			fmt.Println("2", err)
		}
		if err := s2.Shutdown(context.Background()); err != nil {
			fmt.Println("3", err)
		}
		return fmt.Errorf("Recevied signal %s", sig.String())
	case <-ctx.Done():
		return context.Canceled
	}
}
