package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/common/repositories"
	"github.com/tksasha/balance/internal/item"
)

func (r *Repository) FindByID(ctx context.Context, id int) (*item.Item, error) {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    id,
		    date,
		    sum,
		    category_id,
		    category_name,
		    description
		FROM
		    items
		WHERE
		    items.id = ?
		    AND items.deleted_at IS NULL
		    AND currency = ?
	`

	item := &item.Item{}

	if err := r.db.
		QueryRowContext(ctx, query, id, currency).
		Scan(
			&item.ID,
			&item.Date,
			&item.Sum,
			&item.CategoryID,
			&item.CategoryName,
			&item.Description,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrRecordNotFound
		}

		return nil, err
	}

	return item, nil
}
