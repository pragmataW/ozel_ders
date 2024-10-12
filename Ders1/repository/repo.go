package repository

import "errors"

func (r *Repo)GetCredentials(username string) (User, error) {
	if username == "oguz"{
		return User{
			Username: username,
			Password: "123",
		}, nil
	}
	return User{}, errors.New("user not found")
}