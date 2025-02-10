package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/responses"
	"github.com/tksasha/balance/pkg/validation"
)

type CategoryUpdateHandler struct {
	categoryService CategoryService
}

func NewCategoryUpdateHandler(categoryService CategoryService) *CategoryUpdateHandler {
	return &CategoryUpdateHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		var verrors validation.Errors
		if errors.As(err, &verrors) {
			_, _ = w.Write([]byte(verrors.Error()))

			return
		}

		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}

	_ = category

	_, _ = w.Write([]byte("update category page\n"))
}

func (h *CategoryUpdateHandler) handle(r *http.Request) (*models.Category, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, apperrors.ErrResourceNotFound
	}

	category, err := h.categoryService.FindByID(r.Context(), id)
	if err != nil {
		return nil, err
	}

	if err := r.ParseForm(); err != nil {
		return category, apperrors.ErrParsingForm
	}

	category.Name = r.FormValue("name")
	category.Income = r.FormValue("income") == "true"
	category.Visible = r.FormValue("visible") == "true"

	if r.FormValue("supercategory") != "" {
		supercategory, err := strconv.Atoi(r.FormValue("supercategory"))
		if err != nil {
			return nil, apperrors.ErrParsingForm
		}

		category.Supercategory = supercategory
	}

	if err := h.categoryService.Update(r.Context(), category); err != nil {
		return category, err
	}

	return category, nil
}
