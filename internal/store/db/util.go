package db

import "github.com/la4zen/conda-net/internal/model"

func (r *UserDBRepo) Clean(user *model.User) {
	user.Password = ""
}
