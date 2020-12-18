package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/yaoyaoyaoGit/Go-000/Week04/api/demo/v1"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
)

var sigs chan os.Signal

func main() {

	sigs = make(chan os.Signal, 1)
	group, ctx := errgroup.WithContext(context.Background())

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterDemoServer(s, InitializeEvent())

	group.Go(
		func() error {
			if err := s.Serve(lis); err != nil {
				return err
			}
			return nil
		},
	)
	group.Go(
		func() error {
			signal.Notify(sigs)
			select {
			case sig := <-sigs:
				s.GracefulStop()
				return fmt.Errorf("Recevied signal %s", sig.String())
			case <-ctx.Done():
				return context.Canceled
			}
		},
	)

	if err := group.Wait(); err != nil {
		log.Println(err)
	}
}
