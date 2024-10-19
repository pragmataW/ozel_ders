package repository

import (
	"errors"
	"example/dto"
	"example/model"
)

func (r *Repo) GetCredentials(username string) (model.User, error) {
	var user model.User

	result := r.Db.Where("user_name = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		return model.User{}, errors.New("user not found")
	}

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *Repo) AddUser(dtoUser dto.User) error {
	user := model.User{
		Username: dtoUser.UserName,
		Password: dtoUser.Password,
	}

	result := r.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}