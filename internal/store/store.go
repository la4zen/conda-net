package store

import (
	"github.com/la4zen/conda-net/internal/config"
	"github.com/la4zen/conda-net/internal/store/pg"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	DB     *sqlx.DB
	Config *config.Config

	User UserRepo
}

func New() (*Store, error) {
	var err error
	store := &Store{}
	store.Config = config.New()
	store.DB, err = sqlx.Connect("postgres", "")
	if err != nil {
		return nil, err
	}
	store.User = pg.NewUserRepo(store.DB)
	return store, nil
}
