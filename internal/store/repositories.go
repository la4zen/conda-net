package store

import (
	"github.com/la4zen/conda-net/internal/model"
)

type UserRepo interface {
	CreateUser(*model.User) (*model.User, error)
	SetOnlineUser(*model.User) error
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error
	GetUser(*model.User) (*model.User, error)
}
