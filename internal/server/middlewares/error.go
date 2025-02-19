package middlewares

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/response"
)

type errorMiddleware struct{}

func (m *errorMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper, ok := w.(*response.Wrapper)
		if !ok {
			slog.Error("failed to assert wrapper", "wrapper", wrapper)
		}

		next.ServeHTTP(wrapper, r)

		err := wrapper.Error
		if err != nil {
			switch {
			case errors.Is(err, common.ErrParsingForm):
				wrapper.Code = http.StatusBadRequest

				http.Error(wrapper, "Bad Request", wrapper.Code)
			case errors.Is(err, common.ErrResourceNotFound):
				wrapper.Code = http.StatusNotFound

				http.Error(wrapper, "Resource Not Found", wrapper.Code)
			default:
				wrapper.Code = http.StatusInternalServerError

				http.Error(wrapper, "Internal Server Error", wrapper.Code)
			}
		}
	})
}
