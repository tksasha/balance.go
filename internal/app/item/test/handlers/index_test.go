package handlers_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	categoryrepository "github.com/tksasha/balance/internal/app/category/repository"
	"github.com/tksasha/balance/internal/app/item/handlers"
	"github.com/tksasha/balance/internal/app/item/repository"
	"github.com/tksasha/balance/internal/app/item/service"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
)

func TestItemIndexHandler(t *testing.T) {
	handler, db := newIndexHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "GET /items", handler)

	ctx := t.Context()

	t.Run("responds 200 on items found", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/items?currency=eur", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func newIndexHandler(t *testing.T) (*handlers.IndexHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	itemRepository := repository.New(db)

	categoryRepository := categoryrepository.New(db)

	itemService := service.New(itemRepository, categoryRepository)

	handler := handlers.NewIndexHandler(itemService)

	return handler, db
}
