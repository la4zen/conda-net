package store

import (
	"github.com/la4zen/conda-net/internal/model"
)

type UserRepo interface {
	CreateUser(*model.User) *model.Response
	FindUser(*model.User) *model.Response
	GetUser(*model.User) *model.Response
	Clean(*model.User)
	AddFriend(*model.Friend) *model.Response
}
