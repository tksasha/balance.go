package middlewares

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/response"
)

type errorMiddleware struct{}

func newErrorMiddleware() *errorMiddleware {
	return &errorMiddleware{}
}

func (m *errorMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &response.Response{ResponseWriter: w}

		next.ServeHTTP(rw, r)

		if err := rw.Error; err != nil {
			switch {
			case errors.Is(err, common.ErrParsingForm):
				http.Error(w, "Bad Request", http.StatusBadRequest)
			case errors.Is(err, common.ErrResourceNotFound):
				http.Error(w, "Resource Not Found", http.StatusNotFound)
			default:
				slog.Error("Internal Server Error", "error", err)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	})
}
