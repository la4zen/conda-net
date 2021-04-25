package db

import (
	"errors"

	"github.com/la4zen/conda-net/internal/model"
)

func (r *UserDBRepo) AddFriend(friend *model.Friend) *model.Response {
	result := r.db.First(&friend, "from_user = ? or to_user = ?", friend.ToUser, friend.ToUser)
	if result.RowsAffected != 0 {
		return &model.Response{
			Err:  errors.New("user already added"),
			Code: 302,
		}
	}
	r.db.Create(&friend)
	return &model.Response{
		Code: 200,
	}
}
func (r *UserDBRepo) GetFriends(user *model.User) error {
	return nil
}
