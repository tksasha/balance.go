package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/common/handlers"
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common"
)

type UpdateHandler struct {
	service cash.Service
}

func NewUpdateHandler(service cash.Service) *UpdateHandler {
	return &UpdateHandler{
		service: service,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	_, _ = w.Write([]byte(cash.Name))
}

func (h *UpdateHandler) handle(r *http.Request) (*cash.Cash, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := cash.UpdateRequest{
		ID:            r.PathValue("id"),
		Formula:       r.FormValue("formula"),
		Name:          r.FormValue("name"),
		Supercategory: r.FormValue("supercategory"),
		Favorite:      r.FormValue("favorite"),
	}

	return h.service.Update(r.Context(), request)
}
