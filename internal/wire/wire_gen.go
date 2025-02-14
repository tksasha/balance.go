// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/cash/repository"
	"github.com/tksasha/balance/internal/cash/service"
	handlers2 "github.com/tksasha/balance/internal/category/handlers"
	repository2 "github.com/tksasha/balance/internal/category/repository"
	service2 "github.com/tksasha/balance/internal/category/service"
	"github.com/tksasha/balance/internal/config"
	handlers3 "github.com/tksasha/balance/internal/core/index/handler"
	"github.com/tksasha/balance/internal/db"
	handlers4 "github.com/tksasha/balance/internal/item/handlers"
	repository3 "github.com/tksasha/balance/internal/item/repository"
	service3 "github.com/tksasha/balance/internal/item/service"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
)

// Injectors from wire.go:

func InitializeServer() *server.Server {
	configConfig := config.New()
	contextContext := context.Background()
	dbNameProvider := providers.NewDBNameProvider()
	sqlDB := db.Open(contextContext, dbNameProvider)
	repositoryRepository := repository.New(sqlDB)
	serviceService := service.New(repositoryRepository)
	createHandler := handlers.NewCreateHandler(serviceService)
	deleteHandler := handlers.NewDeleteHandler(serviceService)
	editHandler := handlers.NewEditHandler(serviceService)
	listHandler := handlers.NewListHandler(serviceService)
	newHandler := handlers.NewNewHandler()
	updateHandler := handlers.NewUpdateHandler(serviceService)
	repository4 := repository2.New(sqlDB)
	service4 := service2.New(repository4)
	handlersCreateHandler := handlers2.NewCreateHandler(service4)
	handlersDeleteHandler := handlers2.NewDeleteHandler(service4)
	handlersEditHandler := handlers2.NewEditHandler(service4)
	handlersListHandler := handlers2.NewListHandler(service4)
	handlersUpdateHandler := handlers2.NewUpdateHandler(service4)
	handler := handlers3.NewHandler(service4)
	repository5 := repository3.New(sqlDB)
	service5 := service3.New(repository5, repository4)
	createHandler2 := handlers4.NewCreateHandler(service5)
	editHandler2 := handlers4.NewEditHandler(service5)
	listHandler2 := handlers4.NewListHandler(service5)
	updateHandler2 := handlers4.NewUpdateHandler(service5)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, listHandler, newHandler, updateHandler, handlersCreateHandler, handlersDeleteHandler, handlersEditHandler, handlersListHandler, handlersUpdateHandler, handler, createHandler2, editHandler2, listHandler2, updateHandler2)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
