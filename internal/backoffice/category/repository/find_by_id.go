package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) FindByID(ctx context.Context, id int) (*category.Category, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    id,
		    name,
		    income,
		    visible,
		    currency,
		    supercategory
		FROM
		    categories
		WHERE
		    id = ?
		    AND currency = ?
	`

	row := r.db.QueryRowContext(ctx, query, id, currency)

	category := &category.Category{}

	if err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Income,
		&category.Visible,
		&category.Currency,
		&category.Supercategory,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}
