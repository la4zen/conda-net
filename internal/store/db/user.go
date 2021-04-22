package db

import (
	"errors"

	"github.com/la4zen/conda-net/internal/model"
	"gorm.io/gorm"
)

type UserDBRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserDBRepo {
	return &UserDBRepo{db: db}
}

func (r *UserDBRepo) CreateUser(user *model.User) *model.Response {
	defer r.Clean(user)
	result := r.db.First(&user, "login = ?", user.Login)
	if result.RowsAffected != 0 {
		return &model.Response{
			Err:  errors.New("user exists"),
			Code: 401,
		}
	}
	r.db.Create(&user)
	return &model.Response{
		Code: 200,
	}
}

func (r *UserDBRepo) GetUser(user *model.User) *model.Response {
	defer r.Clean(user)
	result := r.db.First(&user, "login = ? or id = ?", user.Login, user.ID)
	if result.RowsAffected != 0 {
		return &model.Response{
			Code: 200,
		}
	}
	return &model.Response{
		Code: 404,
		Err:  errors.New("user not found"),
	}
}

func (r *UserDBRepo) FindUser(user *model.User) *model.Response {
	defer r.Clean(user)
	result := r.db.First(&user, "login = ? and password = ?", user.Login, user.Password)
	if result.RowsAffected != 0 {
		return &model.Response{
			Code: 200,
		}
	}
	return &model.Response{
		Code: 404,
		Err:  errors.New("user not found"),
	}
}

func (r *UserDBRepo) UpdateUser(user *model.User) *model.Response {
	return nil
}

func (r *UserDBRepo) SetOnlineUser(user *model.User) *model.Response {
	return nil
}

func (r *UserDBRepo) Clean(user *model.User) {
	user.Password = ""
}
