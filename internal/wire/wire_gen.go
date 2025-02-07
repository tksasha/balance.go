// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package wire

import (
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
	dbNameProvider := providers.NewDBNameProvider()
	dbDB := db.Open(dbNameProvider)
	cashRepository := repositories.NewCashRepository(dbDB)
	cashService := services.NewCashService(cashRepository)
	cashCreateHandler := handlers.NewCashCreateHandler(cashService)
	cashEditHandler := handlers.NewCashEditHandler(cashService)
	categoryRepository := repositories.NewCategoryRepository(dbDB)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryCreateHandler := handlers.NewCategoryCreateHandler(categoryService)
	categoryDeleteHandler := handlers.NewCategoryDeleteHandler(categoryService)
	categoryEditHandler := handlers.NewCategoryEditHandler(categoryService)
	categoryListHandler := handlers.NewCategoryListHandler(categoryService)
	categoryUpdateHandler := handlers.NewCategoryUpdateHandler(categoryService)
	indexPageHandler := handlers.NewIndexPageHandler(categoryService)
	itemRepository := repositories.NewItemRepository(dbDB)
	itemService := services.NewItemService(itemRepository, categoryRepository)
	itemCreateHandler := handlers.NewItemCreateHandler(itemService)
	itemEditHandler := handlers.NewItemEditHandler(itemService)
	itemListHandler := handlers.NewItemListHandler(itemService)
	routesRoutes := routes.New(cashCreateHandler, cashEditHandler, categoryCreateHandler, categoryDeleteHandler, categoryEditHandler, categoryListHandler, categoryUpdateHandler, indexPageHandler, itemCreateHandler, itemEditHandler, itemListHandler)
	serverServer := server.New(configConfig, routesRoutes)
	return serverServer
}
