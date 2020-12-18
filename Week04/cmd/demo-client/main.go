package main

import (
	"context"
	"log"

	pb "github.com/yaoyaoyaoGit/Go-000/Week04/api/demo/v1"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewDemoClient(conn)
	r, err := c.GetUser(context.Background(), &pb.GetUserRequest{Name: "Test"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.String())
}
