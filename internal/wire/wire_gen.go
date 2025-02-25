// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/app/cash/components"
	handlers2 "github.com/tksasha/balance/internal/app/cash/handlers"
	repository2 "github.com/tksasha/balance/internal/app/cash/repository"
	service2 "github.com/tksasha/balance/internal/app/cash/service"
	repository4 "github.com/tksasha/balance/internal/app/category/repository"
	service4 "github.com/tksasha/balance/internal/app/category/service"
	components2 "github.com/tksasha/balance/internal/app/index/components"
	"github.com/tksasha/balance/internal/app/index/handler"
	repository3 "github.com/tksasha/balance/internal/app/index/repository"
	service3 "github.com/tksasha/balance/internal/app/index/service"
	components3 "github.com/tksasha/balance/internal/app/item/components"
	handlers3 "github.com/tksasha/balance/internal/app/item/handlers"
	repository5 "github.com/tksasha/balance/internal/app/item/repository"
	service5 "github.com/tksasha/balance/internal/app/item/service"
	component2 "github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common/component"
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
	contextContext := context.Background()
	provider := nameprovider.New()
	sqlDB := db.Open(contextContext, provider)
	repositoryRepository := repository.New(sqlDB)
	serviceService := service.New(repositoryRepository)
	componentComponent := component.New()
	categoryComponent := component2.New(componentComponent)
	createHandler := handlers.NewCreateHandler(serviceService, categoryComponent)
	deleteHandler := handlers.NewDeleteHandler(serviceService)
	editHandler := handlers.NewEditHandler(serviceService, categoryComponent)
	listHandler := handlers.NewListHandler(serviceService, categoryComponent)
	updateHandler := handlers.NewUpdateHandler(serviceService, categoryComponent)
	repository6 := repository2.New(sqlDB)
	service6 := service2.New(repository6)
	cashComponent := components.NewCashComponent(componentComponent)
	handlersCreateHandler := handlers2.NewCreateHandler(service6, cashComponent)
	handlersDeleteHandler := handlers2.NewDeleteHandler(service6)
	handlersEditHandler := handlers2.NewEditHandler(service6, cashComponent)
	handlersListHandler := handlers2.NewListHandler(service6, cashComponent)
	newHandler := handlers2.NewNewHandler(cashComponent)
	handlersUpdateHandler := handlers2.NewUpdateHandler(service6, cashComponent)
	repository7 := repository3.New(sqlDB)
	service7 := service3.New(repository7)
	repository8 := repository4.New(sqlDB)
	service8 := service4.New(repository8)
	monthsComponent := components2.NewMonthsComponent(componentComponent)
	yearsComponent := components2.NewYearsComponent(componentComponent)
	indexComponent := components2.NewIndexComponent(componentComponent, monthsComponent, yearsComponent)
	handlerHandler := handler.New(service7, service8, indexComponent)
	repository9 := repository5.New(sqlDB)
	service9 := service5.New(repository9, repository8)
	itemsComponent := components3.NewItemsComponent(componentComponent)
	createHandler2 := handlers3.NewCreateHandler(service9, service8, itemsComponent)
	editHandler2 := handlers3.NewEditHandler(service9, service8, itemsComponent)
	listHandler2 := handlers3.NewListHandler(service9, itemsComponent, monthsComponent, yearsComponent)
	updateHandler2 := handlers3.NewUpdateHandler(service9, service8, itemsComponent)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, listHandler, updateHandler, handlersCreateHandler, handlersDeleteHandler, handlersEditHandler, handlersListHandler, newHandler, handlersUpdateHandler, handlerHandler, createHandler2, editHandler2, listHandler2, updateHandler2)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
