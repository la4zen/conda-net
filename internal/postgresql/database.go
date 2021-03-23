package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/la4zen/conda-net/internal/config"
	"github.com/labstack/gommon/log"
)

type DB struct {
	conn *sqlx.DB
}

func (db *DB) CreateConnection(config config.Config) {
	var err error
	var conString string = fmt.Sprintf("user=%s dbname=%s sslmode=disabled password=%s host=%s",
		config.Database.User, config.Database.DB, config.Database.Password, config.Database.IP)
	db.conn, err = sqlx.Connect("postgres", conString)
	if err != nil {
		log.Error("Database connection error.")
	}
}
