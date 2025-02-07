// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
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
	cashRepository := repositories.NewCashRepository(sqlDB)
	cashService := services.NewCashService(cashRepository)
	cashCreateHandler := handlers.NewCashCreateHandler(cashService)
	cashEditHandler := handlers.NewCashEditHandler(cashService)
	cashDeleteHandler := handlers.NewCashDeleteHandler(cashService)
	categoryRepository := repositories.NewCategoryRepository(sqlDB)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryCreateHandler := handlers.NewCategoryCreateHandler(categoryService)
	categoryDeleteHandler := handlers.NewCategoryDeleteHandler(categoryService)
	categoryEditHandler := handlers.NewCategoryEditHandler(categoryService)
	categoryListHandler := handlers.NewCategoryListHandler(categoryService)
	categoryUpdateHandler := handlers.NewCategoryUpdateHandler(categoryService)
	indexPageHandler := handlers.NewIndexPageHandler(categoryService)
	itemRepository := repositories.NewItemRepository(sqlDB)
	itemService := services.NewItemService(itemRepository, categoryRepository)
	itemCreateHandler := handlers.NewItemCreateHandler(itemService)
	itemEditHandler := handlers.NewItemEditHandler(itemService)
	itemListHandler := handlers.NewItemListHandler(itemService)
	itemUpdateHandler := handlers.NewItemUpdateHandler(itemService)
	routesRoutes := routes.New(cashCreateHandler, cashEditHandler, cashDeleteHandler, categoryCreateHandler, categoryDeleteHandler, categoryEditHandler, categoryListHandler, categoryUpdateHandler, indexPageHandler, itemCreateHandler, itemEditHandler, itemListHandler, itemUpdateHandler)
	serverServer := server.New(configConfig, routesRoutes)
	return serverServer
}
