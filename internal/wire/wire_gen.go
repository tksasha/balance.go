// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/cash/handlers"
	"github.com/tksasha/balance/internal/core/cash/repository"
	"github.com/tksasha/balance/internal/core/cash/service"
	components2 "github.com/tksasha/balance/internal/core/category/components"
	handlers2 "github.com/tksasha/balance/internal/core/category/handlers"
	repository2 "github.com/tksasha/balance/internal/core/category/repository"
	service2 "github.com/tksasha/balance/internal/core/category/service"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/component"
	components3 "github.com/tksasha/balance/internal/core/indexpage/components"
	"github.com/tksasha/balance/internal/core/indexpage/handler"
	repository3 "github.com/tksasha/balance/internal/core/indexpage/repository"
	service3 "github.com/tksasha/balance/internal/core/indexpage/service"
	components4 "github.com/tksasha/balance/internal/core/item/components"
	handlers3 "github.com/tksasha/balance/internal/core/item/handlers"
	repository4 "github.com/tksasha/balance/internal/core/item/repository"
	service4 "github.com/tksasha/balance/internal/core/item/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

// Injectors from wire.go:

func InitializeServer() *server.Server {
	configConfig := config.New()
	baseHandler := common.NewBaseHandler()
	baseRepository := common.NewBaseRepository()
	contextContext := context.Background()
	provider := nameprovider.New()
	sqlDB := db.Open(contextContext, provider)
	repositoryRepository := repository.New(baseRepository, sqlDB)
	serviceService := service.New(repositoryRepository)
	componentComponent := component.New()
	cashComponent := components.NewCashComponent(componentComponent)
	createHandler := handlers.NewCreateHandler(baseHandler, serviceService, cashComponent)
	deleteHandler := handlers.NewDeleteHandler(baseHandler, serviceService)
	editHandler := handlers.NewEditHandler(baseHandler, serviceService, cashComponent)
	listHandler := handlers.NewListHandler(baseHandler, serviceService, cashComponent)
	newHandler := handlers.NewNewHandler(baseHandler, cashComponent)
	updateHandler := handlers.NewUpdateHandler(baseHandler, serviceService, cashComponent)
	baseService := common.NewBaseService()
	repository5 := repository2.New(baseRepository, sqlDB)
	service5 := service2.New(baseService, repository5)
	categoryComponent := components2.NewCategoryComponent(componentComponent)
	handlersCreateHandler := handlers2.NewCreateHandler(baseHandler, service5, categoryComponent)
	handlersDeleteHandler := handlers2.NewDeleteHandler(baseHandler, service5)
	handlersEditHandler := handlers2.NewEditHandler(baseHandler, service5, categoryComponent)
	handlersListHandler := handlers2.NewListHandler(baseHandler, service5, categoryComponent)
	handlersUpdateHandler := handlers2.NewUpdateHandler(baseHandler, service5, categoryComponent)
	repository6 := repository3.New(baseRepository, sqlDB)
	service6 := service3.New(repository6)
	monthsComponent := components3.NewMonthsComponent(componentComponent)
	indexPageComponent := components3.NewIndexPageComponent(componentComponent, monthsComponent)
	handlerHandler := handler.New(baseHandler, service6, service5, indexPageComponent)
	repository7 := repository4.New(baseRepository, sqlDB)
	service7 := service4.New(baseService, repository7, repository5)
	itemsComponent := components4.NewItemsComponent(componentComponent)
	createHandler2 := handlers3.NewCreateHandler(baseHandler, service7, service5, itemsComponent)
	editHandler2 := handlers3.NewEditHandler(baseHandler, service7, service5, itemsComponent)
	listHandler2 := handlers3.NewListHandler(baseHandler, service7, itemsComponent, monthsComponent)
	updateHandler2 := handlers3.NewUpdateHandler(baseHandler, service7, service5, itemsComponent)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, listHandler, newHandler, updateHandler, handlersCreateHandler, handlersDeleteHandler, handlersEditHandler, handlersListHandler, handlersUpdateHandler, handlerHandler, createHandler2, editHandler2, listHandler2, updateHandler2)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
