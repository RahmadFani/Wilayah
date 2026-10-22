package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rahmadfani/wilayah/internal/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewMysql() *bun.DB {
	var sqldb *sql.DB
	var err error

	if config.IS_DEV == "true" {
		sqldb, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME))
	} else {
		sqldb, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME))
	}

	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, mysqldialect.New())

	// Test the connection
	if err := sqldb.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
