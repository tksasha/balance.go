// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package wire

import (
	"context"
	components2 "github.com/tksasha/balance/internal/app/cash/components"
	handlers3 "github.com/tksasha/balance/internal/app/cash/handlers"
	repository3 "github.com/tksasha/balance/internal/app/cash/repository"
	service3 "github.com/tksasha/balance/internal/app/cash/service"
	repository5 "github.com/tksasha/balance/internal/app/category/repository"
	service5 "github.com/tksasha/balance/internal/app/category/service"
	components3 "github.com/tksasha/balance/internal/app/index/components"
	"github.com/tksasha/balance/internal/app/index/handler"
	repository4 "github.com/tksasha/balance/internal/app/index/repository"
	service4 "github.com/tksasha/balance/internal/app/index/service"
	components4 "github.com/tksasha/balance/internal/app/item/components"
	handlers4 "github.com/tksasha/balance/internal/app/item/handlers"
	repository6 "github.com/tksasha/balance/internal/app/item/repository"
	service6 "github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/backoffice/cash/components"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	component2 "github.com/tksasha/balance/internal/backoffice/category/component"
	handlers2 "github.com/tksasha/balance/internal/backoffice/category/handlers"
	repository2 "github.com/tksasha/balance/internal/backoffice/category/repository"
	service2 "github.com/tksasha/balance/internal/backoffice/category/service"
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
	nameProvider := nameprovider.New()
	sqlDB := db.Open(contextContext, nameProvider)
	repositoryRepository := repository.New(sqlDB)
	serviceService := service.New(repositoryRepository)
	cashComponent := components.NewCashComponent()
	createHandler := handlers.NewCreateHandler(serviceService, cashComponent)
	deleteHandler := handlers.NewDeleteHandler(serviceService)
	editHandler := handlers.NewEditHandler(serviceService, cashComponent)
	listHandler := handlers.NewListHandler(serviceService, cashComponent)
	newHandler := handlers.NewNewHandler(cashComponent)
	repository7 := repository2.New(sqlDB)
	service7 := service2.New(repository7)
	componentComponent := component.New()
	categoryComponent := component2.New(componentComponent)
	handlersCreateHandler := handlers2.NewCreateHandler(service7, categoryComponent)
	handlersDeleteHandler := handlers2.NewDeleteHandler(service7)
	handlersEditHandler := handlers2.NewEditHandler(service7, categoryComponent)
	handlersListHandler := handlers2.NewListHandler(service7, categoryComponent)
	updateHandler := handlers2.NewUpdateHandler(service7, categoryComponent)
	repository8 := repository3.New(sqlDB)
	service8 := service3.New(repository8)
	componentsCashComponent := components2.NewCashComponent()
	editHandler2 := handlers3.NewEditHandler(service8, componentsCashComponent)
	handlersUpdateHandler := handlers3.NewUpdateHandler(service8, componentsCashComponent)
	repository9 := repository4.New(sqlDB)
	service9 := service4.New(repository9)
	repository10 := repository5.New(sqlDB)
	service10 := service5.New(repository10)
	monthsComponent := components3.NewMonthsComponent(componentComponent)
	yearsComponent := components3.NewYearsComponent(componentComponent)
	indexComponent := components3.NewIndexComponent(componentComponent, monthsComponent, yearsComponent)
	handlerHandler := handler.New(service9, service10, indexComponent)
	repository11 := repository6.New(sqlDB)
	service11 := service6.New(repository11, repository10)
	itemsComponent := components4.NewItemsComponent(componentComponent)
	createHandler2 := handlers4.NewCreateHandler(service11, service10, itemsComponent)
	editHandler3 := handlers4.NewEditHandler(service11, service10, itemsComponent)
	listHandler2 := handlers4.NewListHandler(service11, itemsComponent, monthsComponent, yearsComponent)
	updateHandler2 := handlers4.NewUpdateHandler(service11, service10, itemsComponent)
	routesRoutes := routes.New(createHandler, deleteHandler, editHandler, listHandler, newHandler, handlersCreateHandler, handlersDeleteHandler, handlersEditHandler, handlersListHandler, updateHandler, editHandler2, handlersUpdateHandler, handlerHandler, createHandler2, editHandler3, listHandler2, updateHandler2)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
