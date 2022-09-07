//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-crsdemo/app/order/internal/biz"
	"kratos-crsdemo/app/order/internal/conf"
	"kratos-crsdemo/app/order/internal/data"
	"kratos-crsdemo/app/order/internal/server"
	"kratos-crsdemo/app/order/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
