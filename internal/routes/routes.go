package routes

import (
	"embed"
	"net/http"

	category "github.com/tksasha/balance/internal/category/handlers"
	cash "github.com/tksasha/balance/internal/core/cash/handlers"
	index "github.com/tksasha/balance/internal/core/index/handler"
	item "github.com/tksasha/balance/internal/core/item/handlers"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New(
	cashCreateHandler *cash.CreateHandler,
	cashDeleteHandler *cash.DeleteHandler,
	cashEditHandler *cash.EditHandler,
	cashListHandler *cash.ListHandler,
	cashNewHandler *cash.NewHandler,
	cashUpdateHandler *cash.UpdateHandler,
	categoryCreateHandler *category.CreateHandler,
	categoryDeleteHandler *category.DeleteHandler,
	categoryEditHandler *category.EditHandler,
	categoryListHandler *category.ListHandler,
	categoryUpdateHandler *category.UpdateHandler,
	indexPageHandler *index.Handler,
	itemCreateHandler *item.CreateHandler,
	itemEditHandler *item.EditHandler,
	itemListHandler *item.ListHandler,
	itemUpdateHandler *item.UpdateHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", indexPageHandler)

	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items", itemListHandler)
	mux.Handle("GET /items/{id}/edit", itemEditHandler)
	mux.Handle("PATCH /items/{id}", itemUpdateHandler)

	mux.Handle("POST /categories", categoryCreateHandler)
	mux.Handle("GET /categories", categoryListHandler)
	mux.Handle("GET /categories/{id}/edit", categoryEditHandler)
	mux.Handle("PATCH /categories/{id}", categoryUpdateHandler)
	mux.Handle("DELETE /categories/{id}", categoryDeleteHandler)

	mux.Handle("DELETE /cashes/{id}", cashDeleteHandler)
	mux.Handle("GET /cashes", cashListHandler)
	mux.Handle("GET /cashes/new", cashNewHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)
	mux.Handle("PATCH /cashes/{id}", cashUpdateHandler)
	mux.Handle("POST /cash", cashCreateHandler)

	return &Routes{
		Mux: mux,
	}
}
