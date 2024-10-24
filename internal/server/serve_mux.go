package server

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
)

//go:embed assets
var assets embed.FS

func NewServeMux(routes []handlers.Route) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}
