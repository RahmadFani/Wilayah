package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/rahmadfani/wilayah/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func Migrate(db *bun.DB) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	migrator := migrate.NewMigrator(db, migrations.Migrations)

	migrator.Init(ctx)

	if err := migrator.Lock(ctx); err != nil {
		slog.Error("migrator.Lock", "error_message", err.Error())
		return
	}
	defer migrator.Unlock(ctx) //nolint:errcheck

	group, err := migrator.Migrate(ctx)
	if err != nil {
		slog.Error("migrator.Migrate", "error_message", err.Error())
		return
	}
	if group.IsZero() {
		slog.Info("there are no new migrations to run (database is up to date)\n")
		return
	}
	fmt.Printf("migrated to %s\n", group)

}
