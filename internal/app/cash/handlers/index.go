package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*commonhandler.Handler

	service   cash.Service
	component *component.Component
}

func NewIndexHandler(service cash.Service) *IndexHandler {
	return &IndexHandler{
		Handler:   commonhandler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.service.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.Index(cashes).Render(w)

	h.SetError(w, err)
}
