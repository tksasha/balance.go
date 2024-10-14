package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tksasha/balance/internal/server/db/migrations"
	"github.com/tksasha/balance/internal/server/env"
	"github.com/tksasha/balance/internal/server/workdir"
)

func Open() *sql.DB {
	ctx := context.Background()

	dbName := fmt.Sprintf(
		"%s%s%s.sqlite3",
		workdir.New(),
		string(os.PathSeparator),
		env.Get(),
	)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			panic(err)
		}

		panic(err)
	}

	if _, err := db.ExecContext(ctx, "PRAGMA foreign_keys=true"); err != nil {
		panic(err)
	}

	migrate(ctx, db)

	return db
}

func migrate(ctx context.Context, db *sql.DB) {
	migrations.
		NewAddItemsCategoryNameMigration(db).Up(ctx)
}
