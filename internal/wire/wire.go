//go:build wireinject

package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/tksasha/balance/internal/app/cash"
	cashcomponents "github.com/tksasha/balance/internal/app/cash/components"
	cashhandlers "github.com/tksasha/balance/internal/app/cash/handlers"
	cashrepository "github.com/tksasha/balance/internal/app/cash/repository"
	cashservice "github.com/tksasha/balance/internal/app/cash/service"
	"github.com/tksasha/balance/internal/app/category"
	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	categoryservice "github.com/tksasha/balance/internal/app/category/service"
	"github.com/tksasha/balance/internal/app/index"
	indexcomponents "github.com/tksasha/balance/internal/app/index/components"
	indexhandler "github.com/tksasha/balance/internal/app/index/handler"
	indexrepository "github.com/tksasha/balance/internal/app/index/repository"
	indexservice "github.com/tksasha/balance/internal/app/index/service"
	"github.com/tksasha/balance/internal/app/item"
	itemcomponents "github.com/tksasha/balance/internal/app/item/components"
	itemhandlers "github.com/tksasha/balance/internal/app/item/handlers"
	itemrepository "github.com/tksasha/balance/internal/app/item/repository"
	itemservice "github.com/tksasha/balance/internal/app/item/service"
	backofficecash "github.com/tksasha/balance/internal/backoffice/cash"
	backofficecashcomponents "github.com/tksasha/balance/internal/backoffice/cash/components"
	backofficecashhandlers "github.com/tksasha/balance/internal/backoffice/cash/handlers"
	backofficecashrepository "github.com/tksasha/balance/internal/backoffice/cash/repository"
	backofficecashservice "github.com/tksasha/balance/internal/backoffice/cash/service"
	backofficecategory "github.com/tksasha/balance/internal/backoffice/category"
	backofficecategorycomponent "github.com/tksasha/balance/internal/backoffice/category/component"
	backofficecategoryhandlers "github.com/tksasha/balance/internal/backoffice/category/handlers"
	backofficecategoryrepository "github.com/tksasha/balance/internal/backoffice/category/repository"
	backofficecategoryservice "github.com/tksasha/balance/internal/backoffice/category/service"
	"github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"github.com/tksasha/balance/internal/server"
	"github.com/tksasha/balance/internal/server/config"
	"github.com/tksasha/balance/internal/server/middlewares"
	"github.com/tksasha/balance/internal/server/routes"
)

func InitializeServer() *server.Server {
	wire.Build(
		backofficecashcomponents.NewCashComponent,
		backofficecashhandlers.NewCreateHandler,
		backofficecashrepository.New,
		backofficecashservice.New,
		backofficecategorycomponent.New,
		backofficecategoryhandlers.NewCreateHandler,
		backofficecategoryhandlers.NewDeleteHandler,
		backofficecategoryhandlers.NewEditHandler,
		backofficecategoryhandlers.NewListHandler,
		backofficecategoryhandlers.NewUpdateHandler,
		backofficecategoryrepository.New,
		backofficecategoryservice.New,
		cashcomponents.NewCashComponent,
		cashhandlers.NewDeleteHandler,
		cashhandlers.NewEditHandler,
		cashhandlers.NewListHandler,
		cashhandlers.NewNewHandler,
		cashhandlers.NewUpdateHandler,
		cashrepository.New,
		cashservice.New,
		categoryrepository.New,
		categoryservice.New,
		component.New,
		config.New,
		context.Background,
		db.Open,
		indexcomponents.NewIndexComponent,
		indexcomponents.NewMonthsComponent,
		indexcomponents.NewYearsComponent,
		indexhandler.New,
		indexrepository.New,
		indexservice.New,
		itemcomponents.NewItemsComponent,
		itemhandlers.NewCreateHandler,
		itemhandlers.NewEditHandler,
		itemhandlers.NewListHandler,
		itemhandlers.NewUpdateHandler,
		itemrepository.New,
		itemservice.New,
		middlewares.New,
		nameprovider.New,
		routes.New,
		server.New,
		wire.Bind(new(backofficecash.Repository), new(*backofficecashrepository.Repository)),
		wire.Bind(new(backofficecash.Service), new(*backofficecashservice.Service)),
		wire.Bind(new(backofficecategory.Repository), new(*backofficecategoryrepository.Repository)),
		wire.Bind(new(backofficecategory.Service), new(*backofficecategoryservice.Service)),
		wire.Bind(new(cash.Repository), new(*cashrepository.Repository)),
		wire.Bind(new(cash.Service), new(*cashservice.Service)),
		wire.Bind(new(category.Repository), new(*categoryrepository.Repository)),
		wire.Bind(new(db.NameProvider), new(*nameprovider.NameProvider)),
		wire.Bind(new(index.CategoryService), new(*categoryservice.Service)),
		wire.Bind(new(index.Repository), new(*indexrepository.Repository)),
		wire.Bind(new(index.Service), new(*indexservice.Service)),
		wire.Bind(new(item.CategoryRepository), new(*categoryrepository.Repository)),
		wire.Bind(new(item.CategoryService), new(*categoryservice.Service)),
		wire.Bind(new(item.Repository), new(*itemrepository.Repository)),
		wire.Bind(new(item.Service), new(*itemservice.Service)),
	)

	return nil
}
