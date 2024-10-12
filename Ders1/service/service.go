package service

import (
	"errors"
	"example/dto"
)

func (s *Service) Login(dtoUser dto.User) (string, error) {
	user, err := s.repo.GetCredentials(dtoUser.UserName)
	
	if err != nil{
		return "", errors.New("db error")
	}

	if user.Password != dtoUser.Password{
		return "", errors.New("invalid credentials")
	}

	return "abc123", nil
}