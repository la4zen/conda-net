package store

import (
	"github.com/la4zen/conda-net/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	DB     *sqlx.DB
	Config *config.Config
}

func New() (*Store, error) {
	var err error
	store := &Store{}
	store.Config = config.New()
	store.DB, err = sqlx.Connect("postgres", "")
	if err != nil {
		return nil, err
	}
	return store, nil
}
