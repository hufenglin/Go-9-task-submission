//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/starryrbs/kfan/app/history/service/internal/biz"
	"github.com/starryrbs/kfan/app/history/service/internal/conf"
	"github.com/starryrbs/kfan/app/history/service/internal/data"
	"github.com/starryrbs/kfan/app/history/service/internal/server"
	"github.com/starryrbs/kfan/app/history/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
