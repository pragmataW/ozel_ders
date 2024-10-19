package repository

import (
	"errors"
	"example/model"
)

func (r *Repo) GetCredentials(username string) (model.User, error) {
	var user model.User

	result := r.Db.Where("username = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		return model.User{}, errors.New("user not found")
	}

	if result.Error != nil{
		return model.User{}, result.Error
	}

	return user, nil
}
