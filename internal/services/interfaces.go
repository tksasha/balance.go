package services

import (
	"context"

	"github.com/tksasha/balance/internal/models"
)

type ItemRepository interface {
	GetItems(ctx context.Context) (models.Items, error)
	Create(ctx context.Context, item *models.Item) error
	FindByID(ctx context.Context, id int) (*models.Item, error)
	Update(ctx context.Context, item *models.Item) error
	Delete(ctx context.Context, id int) error
}

type CategoryRepository interface {
	GetAll(ctx context.Context) (models.Categories, error)
	Create(ctx context.Context, category *models.Category) error
	FindByName(ctx context.Context, name string) (*models.Category, error)
	FindByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, category *models.Category) error
}

type CashRepository interface {
	Create(ctx context.Context, cash *models.Cash) error
	NameExists(ctx context.Context, name string, id int) (bool, error)
	FindByID(ctx context.Context, id int) (*models.Cash, error)
	Update(ctx context.Context, cash *models.Cash) error
}
