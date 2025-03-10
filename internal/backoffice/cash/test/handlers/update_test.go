package handlers_test

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/backoffice/cash/handlers"
	"github.com/tksasha/balance/internal/backoffice/cash/repository"
	"github.com/tksasha/balance/internal/backoffice/cash/service"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/db/nameprovider"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestCashUpdateHandler(t *testing.T) { //nolint:funlen
	handler, db := newUpdateHandler(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	mux := mux(t, "PATCH /backoffice/cashes/{id}", handler)

	ctx := t.Context()

	t.Run("renders 404 when cash not found", func(t *testing.T) {
		formData := url.Values{}

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPatch,
			"/backoffice/cashes/1439",
			strings.NewReader(formData.Encode()),
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-form-www-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNotFound)
	})

	t.Run("renders 400 when invalid input", func(t *testing.T) {
		request, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/backoffice/cashes/1453", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusBadRequest)
	})

	t.Run("renders errors when validation failed", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:       1418,
			Currency: currency.USD,
		}

		createCash(t, db, cashToCreate)

		values := url.Values{
			"name": {""},
		}

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPatch,
			"/backoffice/cashes/1418?currency=usd",
			strings.NewReader(values.Encode()),
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		response, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatal(err)
		}

		golden.Assert(t, string(response), "update-with-errors.html")
	})

	t.Run("renders 204 when cash updated", func(t *testing.T) {
		cleanup(t, db)

		cashToCreate := &cash.Cash{
			ID:            1442,
			Currency:      currency.UAH,
			Formula:       "2+3",
			Sum:           5,
			Name:          "Bonds",
			Supercategory: 2,
			Favorite:      false,
		}

		createCash(t, db, cashToCreate)

		values := url.Values{
			"formula":       {"3+4"},
			"name":          {"Stocks"},
			"supercategory": {"3"},
			"favorite":      {"true"},
		}

		request, err := http.NewRequestWithContext(
			ctx,
			http.MethodPatch,
			"/backoffice/cashes/1442",
			strings.NewReader(values.Encode()),
		)
		if err != nil {
			t.Fatal(err)
		}

		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusNoContent)
	})
}

func newUpdateHandler(t *testing.T) (*handlers.UpdateHandler, *sql.DB) {
	t.Helper()

	db := db.Open(t.Context(), nameprovider.NewTestProvider())

	cashRepository := repository.New(db)

	cashService := service.New(cashRepository)

	handler := handlers.NewUpdateHandler(cashService)

	return handler, db
}
