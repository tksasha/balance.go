package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/tests"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/handlers"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestDeleteItemHandler(t *testing.T) {
	ctx := t.Context()

	service, db := tests.NewItemService(ctx, t)
	defer func() {
		_ = db.Close()
	}()

	mux := tests.NewMux(t, "DELETE /items/{id}", handlers.NewDeleteHandler(service))

	t.Run("renders 404 when item is not found", func(t *testing.T) {
		request := tests.NewDeleteRequest(ctx, t, "/items/1043")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("render 204 when item is deleted", func(t *testing.T) {
		tests.Cleanup(ctx, t)

		categoryToCreate := &category.Category{
			ID:       1047,
			Currency: currencies.UAH,
		}

		tests.CreateCategory(ctx, t, categoryToCreate)

		itemToCreate := &item.Item{
			ID:         1045,
			Currency:   currencies.UAH,
			CategoryID: 1047,
		}

		tests.CreateItem(ctx, t, itemToCreate)

		request := tests.NewDeleteRequest(ctx, t, "/items/1045")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)
	})
}
