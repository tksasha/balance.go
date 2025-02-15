package service

import (
	"context"

	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/month"
)

func (s *Service) List(ctx context.Context, request item.IndexRequest) (item.Items, error) {
	month := month.New(request.Year, request.Month)

	return s.itemRepository.List(ctx, month)
}
