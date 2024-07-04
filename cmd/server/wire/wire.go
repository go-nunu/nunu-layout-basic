//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-basic/internal/handler"
	"github.com/go-nunu/nunu-layout-basic/internal/repository"
	"github.com/go-nunu/nunu-layout-basic/internal/server"
	"github.com/go-nunu/nunu-layout-basic/internal/service"
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func NewWire(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
