package server

import (
	"context"
	"fmt"

	pb "github.com/yaoyaoyaoGit/Go-000/Week04/api/demo/v1"
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/service"
)

type (
	DemoServer struct {
		pb.UnimplementedDemoServer
		UserService *service.UserService
	}
)

func (s *DemoServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.UserService.GetUser(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("user not found")
	}
	return &pb.GetUserResponse{
		Name:  u.UserName,
		Age:   u.Age,
		Email: u.Email,
	}, nil
}

func NewDemoServer(s *service.UserService) *DemoServer {
	demo := &DemoServer{
		UserService: s,
	}
	return demo
}
