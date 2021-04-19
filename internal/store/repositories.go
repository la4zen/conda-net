package store

import (
	"github.com/la4zen/conda-net/internal/model"
)

type UserRepo interface {
	CreateUser(*model.User) (*model.User, int, error)
	SetOnlineUser(*model.User) (int, error)
	UpdateUser(*model.User) (*model.User, int, error)
	DeleteUser(*model.User) (int, error)
	GetUser(*model.User) (*model.User, int, error)
}
