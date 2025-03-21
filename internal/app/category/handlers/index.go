package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/category/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
)

type IndexHandler struct {
	*commonhandler.Handler

	service   category.Service
	component *component.Component
}

func NewIndexHandler(service category.Service) *IndexHandler {
	return &IndexHandler{
		Handler:   commonhandler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := category.Request{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	entities, err := h.service.GroupedList(r.Context(), request)
	if err != nil {
		h.SetError(w, err)
	}

	params := params.New(r.URL.Query())

	err = h.component.Index(entities, params).Render(w)

	h.SetError(w, err)
}
