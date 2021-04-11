package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/la4zen/conda-net/internal/model"
)

type UserPgRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserPgRepo {
	return &UserPgRepo{db: db}
}

func (repo *UserPgRepo) CreateUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (repo *UserPgRepo) GetUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (repo *UserPgRepo) UpdateUser(user *model.User) (*model.User, error) {
	return nil, nil
}

func (repo *UserPgRepo) DeleteUser(user *model.User) error {
	return nil
}

func (repo *UserPgRepo) SetOnlineUser(user *model.User) error {
	return nil
}
