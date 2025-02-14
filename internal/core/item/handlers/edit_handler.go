package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type EditHandler struct {
	service item.Service
}

func NewEditHandler(service item.Service) *EditHandler {
	return &EditHandler{
		service: service,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(w, r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Edit(item).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *EditHandler) handle(_ http.ResponseWriter, r *http.Request) (*item.Item, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
