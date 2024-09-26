package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/server/app"
)

type GetItemsHandler struct {
	itemRepository *repositories.ItemRepository
}

func NewGetItemsHandler(app *app.App) http.Handler {
	return &GetItemsHandler{
		itemRepository: repositories.NewItemRepository(app.DB),
	}
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.itemRepository.GetItems(r.Context())
	if err != nil {
		slog.Error(err.Error())
	}

	if err := components.ItemList(items).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
