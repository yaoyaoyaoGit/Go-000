// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/model"
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/server"
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/service"
)

// Injectors from wire.go:

func InitializeEvent() *server.DemoServer {
	codeModel := model.NewCodeModel()
	userService := service.NewUserService(codeModel)
	demoServer := server.NewDemoServer(userService)
	return demoServer
}
