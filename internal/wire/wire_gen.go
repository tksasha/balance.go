// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/cash/repository"
	"github.com/tksasha/balance/internal/cash/service"
	handlers3 "github.com/tksasha/balance/internal/category/handlers"
	repository2 "github.com/tksasha/balance/internal/category/repository"
	service2 "github.com/tksasha/balance/internal/category/service"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	handlers2 "github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/middlewares"
	"github.com/tksasha/balance/internal/providers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/services"
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
	categoryRepository := repositories.NewCategoryRepository(sqlDB)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryCreateHandler := handlers2.NewCategoryCreateHandler(categoryService)
	repository3 := repository2.New(sqlDB)
	service3 := service2.New(repository3)
	handlersDeleteHandler := handlers3.NewDeleteHandler(service3)
	categoryEditHandler := handlers2.NewCategoryEditHandler(categoryService)
	handlersListHandler := handlers3.NewListHandler(service3)
	categoryUpdateHandler := handlers2.NewCategoryUpdateHandler(categoryService)
	indexPageHandler := handlers2.NewIndexPageHandler(service3)
	itemRepository := repositories.NewItemRepository(sqlDB)
	itemService := services.NewItemService(itemRepository, categoryRepository)
	itemCreateHandler := handlers2.NewItemCreateHandler(itemService)
	itemEditHandler := handlers2.NewItemEditHandler(itemService)
	itemListHandler := handlers2.NewItemListHandler(itemService)
	itemUpdateHandler := handlers2.NewItemUpdateHandler(itemService)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, listHandler, newHandler, updateHandler, categoryCreateHandler, handlersDeleteHandler, categoryEditHandler, handlersListHandler, categoryUpdateHandler, indexPageHandler, itemCreateHandler, itemEditHandler, itemListHandler, itemUpdateHandler)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
