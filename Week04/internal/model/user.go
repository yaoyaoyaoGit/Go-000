package model

import (
	"context"
	"fmt"
)

type User struct {
	InternalID string
	UserName   string
	Password   string
	Age        string
	Email      string
}

type CodeModel struct{}

func (*CodeModel) GetUser(ctx context.Context, UserName string) (*User, error) {
	if UserName == "Test" {
		return &User{
			InternalID: "1",
			UserName:   "Test",
			Password:   "Hash",
			Age:        "42",
			Email:      "TestEmail@gmail.com",
		}, nil
	}
	return nil, fmt.Errorf("user not found")
}

func NewCodeModel() *CodeModel {
	return &CodeModel{}
}
