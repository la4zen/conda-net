package store

import (
	"github.com/la4zen/conda-net/internal/config"
	"github.com/la4zen/conda-net/internal/model"
	"github.com/la4zen/conda-net/internal/store/db"
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
	store.DB, err = gorm.Open(postgres.Open("host=185.199.9.125 user=postgres password=897+897 database=conda sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	store.DB.AutoMigrate(&model.User{}, &model.Friend{})
	store.User = db.NewUserRepo(store.DB)
	return store, nil
}
