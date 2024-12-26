// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package wire

import (
	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/routes"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/services"
)

// Injectors from wire.go:

func InitializeServer() *server.Server {
	configConfig := config.New()
	dbDB := db.Open()
	categoryRepository := repositories.NewCategoryRepository(dbDB)
	categoryService := services.NewCategoryService(categoryRepository)
	indexPageHandler := handlers.NewIndexPageHandler(categoryService)
	itemRepository := repositories.NewItemRepository(dbDB)
	itemService := services.NewItemService(itemRepository)
	createItemHandler := handlers.NewCreateItemHandler(itemService, categoryService)
	getItemsHandler := handlers.NewGetItemsHandler(itemService)
	getItemHandler := handlers.NewGetItemHandler(itemService)
	getCategoriesHandler := handlers.NewGetCategoriesHandler(categoryService)
	routesRoutes := routes.New(indexPageHandler, createItemHandler, getItemsHandler, getItemHandler, getCategoriesHandler)
	serverServer := server.New(configConfig, routesRoutes)
	return serverServer
}
