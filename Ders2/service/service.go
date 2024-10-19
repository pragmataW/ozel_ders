package service

import (
	"errors"
	"example/dto"
	"log"
)

func (s *Service) Login(dtoUser dto.User) (string, error) {
	user, err := s.repo.GetCredentials(dtoUser.UserName)
	
	if err != nil{
		return "", err
	}

	decryptedPass, err := s.encryptor.Decrypt(user.Password)
	if err != nil{
		log.Println("decrypt error")
		return "", err
	}

	if dtoUser.Password != decryptedPass{
		return "", errors.New("invalid credentials")
	}

	jwtKey, err := s.jwtGenerator.Generate(dtoUser.UserName)
	if err != nil{
		log.Println("jwt generator err")
		return "", err
	}

	return jwtKey, nil
}

func (s *Service) Register(user dto.User) (error) {
	_, err := s.repo.GetCredentials(user.UserName)
	if err.Error() != "user not found"{
		return errors.New("user already exists")
	}
	
	encryptedPassword, err := s.encryptor.Encrypt(user.Password)
	if err != nil{
		return err
	}
	user.Password = encryptedPassword

	if err := s.repo.AddUser(user); err != nil{
		return err
	}

	return nil
}