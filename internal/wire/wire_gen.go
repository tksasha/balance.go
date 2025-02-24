// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/tksasha/balance/internal/app/cash/components"
	handlers2 "github.com/tksasha/balance/internal/app/cash/handlers"
	repository2 "github.com/tksasha/balance/internal/app/cash/repository"
	service2 "github.com/tksasha/balance/internal/app/cash/service"
	components2 "github.com/tksasha/balance/internal/app/category/components"
	handlers3 "github.com/tksasha/balance/internal/app/category/handlers"
	repository3 "github.com/tksasha/balance/internal/app/category/repository"
	service3 "github.com/tksasha/balance/internal/app/category/service"
	components3 "github.com/tksasha/balance/internal/app/index/components"
	"github.com/tksasha/balance/internal/app/index/handler"
	repository4 "github.com/tksasha/balance/internal/app/index/repository"
	service4 "github.com/tksasha/balance/internal/app/index/service"
	components4 "github.com/tksasha/balance/internal/app/item/components"
	handlers4 "github.com/tksasha/balance/internal/app/item/handlers"
	repository5 "github.com/tksasha/balance/internal/app/item/repository"
	service5 "github.com/tksasha/balance/internal/app/item/service"
	component2 "github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/backoffice/category/handlers"
	"github.com/tksasha/balance/internal/backoffice/category/repository"
	"github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common"
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
	baseHandler := common.NewBaseHandler()
	baseService := common.NewBaseService()
	baseRepository := common.NewBaseRepository()
	contextContext := context.Background()
	provider := nameprovider.New()
	sqlDB := db.Open(contextContext, provider)
	repositoryRepository := repository.New(baseRepository, sqlDB)
	serviceService := service.New(baseService, repositoryRepository)
	componentComponent := component.New()
	categoryComponent := component2.New(componentComponent)
	createHandler := handlers.NewCreateHandler(baseHandler, serviceService, categoryComponent)
	deleteHandler := handlers.NewDeleteHandler(baseHandler, serviceService)
	listHandler := handlers.NewListHandler(baseHandler, serviceService, categoryComponent)
	repository6 := repository2.New(baseRepository, sqlDB)
	service6 := service2.New(repository6)
	cashComponent := components.NewCashComponent(componentComponent)
	handlersCreateHandler := handlers2.NewCreateHandler(baseHandler, service6, cashComponent)
	handlersDeleteHandler := handlers2.NewDeleteHandler(baseHandler, service6)
	editHandler := handlers2.NewEditHandler(baseHandler, service6, cashComponent)
	handlersListHandler := handlers2.NewListHandler(baseHandler, service6, cashComponent)
	newHandler := handlers2.NewNewHandler(baseHandler, cashComponent)
	updateHandler := handlers2.NewUpdateHandler(baseHandler, service6, cashComponent)
	repository7 := repository3.New(baseRepository, sqlDB)
	service7 := service3.New(baseService, repository7)
	componentsCategoryComponent := components2.NewCategoryComponent(componentComponent)
	handlersEditHandler := handlers3.NewEditHandler(baseHandler, service7, componentsCategoryComponent)
	listHandler2 := handlers3.NewListHandler(baseHandler, service7, componentsCategoryComponent)
	handlersUpdateHandler := handlers3.NewUpdateHandler(baseHandler, service7, componentsCategoryComponent)
	repository8 := repository4.New(baseRepository, sqlDB)
	service8 := service4.New(repository8)
	monthsComponent := components3.NewMonthsComponent(componentComponent)
	yearsComponent := components3.NewYearsComponent(componentComponent)
	indexComponent := components3.NewIndexComponent(componentComponent, monthsComponent, yearsComponent)
	handlerHandler := handler.New(baseHandler, service8, service7, indexComponent)
	repository9 := repository5.New(baseRepository, sqlDB)
	service9 := service5.New(baseService, repository9, repository7)
	itemsComponent := components4.NewItemsComponent(componentComponent)
	createHandler2 := handlers4.NewCreateHandler(baseHandler, service9, service7, itemsComponent)
	editHandler2 := handlers4.NewEditHandler(baseHandler, service9, service7, itemsComponent)
	listHandler3 := handlers4.NewListHandler(baseHandler, service9, itemsComponent, monthsComponent, yearsComponent)
	updateHandler2 := handlers4.NewUpdateHandler(baseHandler, service9, service7, itemsComponent)
	routesRoutes := routes.New(createHandler, deleteHandler, listHandler, handlersCreateHandler, handlersDeleteHandler, editHandler, handlersListHandler, newHandler, updateHandler, handlersEditHandler, listHandler2, handlersUpdateHandler, handlerHandler, createHandler2, editHandler2, listHandler3, updateHandler2)
	v := middlewares.New()
	serverServer := server.New(configConfig, routesRoutes, v)
	return serverServer
}
