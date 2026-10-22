package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/rahmadfani/wilayah/internal/config"
	"github.com/rahmadfani/wilayah/internal/database"
	"github.com/rahmadfani/wilayah/internal/handler"
)

func main() {

	config.New()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	db := database.NewMysql()
	defer db.Close()

	app := handler.NewHttp(db)

	log.Fatal(app.Listen(":9876"))
}
