package store

import (
	"github.com/la4zen/conda-net/internal/config"
	"github.com/la4zen/conda-net/internal/model"
	"github.com/la4zen/conda-net/internal/store/pg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Store struct {
	DB     *gorm.DB
	Config *config.Config

	User UserRepo
}

func New() (*Store, error) {
	var err error
	store := &Store{}
	store.Config = config.New()
	store.DB, err = gorm.Open(postgres.Open("user=postgres password=897+897 database=postgres sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	store.DB.AutoMigrate(&model.User{})
	store.User = pg.NewUserRepo(store.DB)
	return store, nil
}
