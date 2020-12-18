//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/model"
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/server"
	"github.com/yaoyaoyaoGit/Go-000/Week04/internal/service"
)

func InitializeEvent() *server.DemoServer {
	wire.Build(model.NewCodeModel,
		wire.Bind(new(service.UserGetter), new(*model.CodeModel)),
		service.NewUserService,
		server.NewDemoServer,
	)
	return &server.DemoServer{}
}
