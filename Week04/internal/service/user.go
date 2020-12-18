package service

import (
	"context"

	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/model"
)

type (
	UserGetter interface {
		GetUser(ctx context.Context, UserName string) (*model.User, error)
	}
	UserService struct {
		m UserGetter
	}
)

func (b *UserService) GetUser(ctx context.Context, name string) (*model.User, error) {
	u, err := b.m.GetUser(ctx, name)
	if err != nil {
		return nil, err
	}

	// some business logic here
	return u, nil
}

func NewUserService(getter UserGetter) *UserService {
	return &UserService{getter}
}
