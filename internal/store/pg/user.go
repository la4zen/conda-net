package pg

import (
	"errors"

	"github.com/la4zen/conda-net/internal/model"
	"gorm.io/gorm"
)

type UserPgRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserPgRepo {
	return &UserPgRepo{db: db}
}

func (r *UserPgRepo) CreateUser(user *model.User) (*model.User, int, error) {
	user, code, err := r.GetUser(user)
	if code == 200 {
		return nil, 401, errors.New("User already exists!")
	}
	if err != nil && code == 500 {
		return nil, code, err
	}
	result := r.db.Create(&user)
	return user, code, result.Error
}

func (r *UserPgRepo) GetUser(user *model.User) (*model.User, int, error) {
	result := r.db.Find(&model.User{}, &user)
	if result.Error != nil {
		return nil, 500, result.Error
	}
	if result.RowsAffected == 0 {
		return user, 404, errors.New("User not found")
	}
	return user, 200, nil
}

func (r *UserPgRepo) UpdateUser(user *model.User) (*model.User, int, error) {
	return nil, 200, nil
}

func (r *UserPgRepo) DeleteUser(user *model.User) (int, error) {
	return 200, nil
}

func (r *UserPgRepo) SetOnlineUser(user *model.User) (int, error) {
	return 200, nil
}
