package service

import (
	"example/dto"
	"example/model"
	mocks "example/service_mocks"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	authRepo := &mocks.IAuthRepo{}
    jwtGenerator := &mocks.IJwtGenerator{}
    encryptor := &mocks.IEncryptor{}

	user := dto.User{
		UserName: "oguz",
		Password: "123wsedrf",
	}
	key := "my_ultra_jwt_key"
	hashedPass := "123"

	modelUser := model.User {
		Username: "oguz",
		Password: hashedPass,
	}
	
    authRepo.On("GetCredentials", user.UserName).Return(modelUser, nil)
	jwtGenerator.On("Generate", user.UserName).Return(key, nil)
	encryptor.On("Decrypt", hashedPass).Return(user.Password, nil)
	
	srv := New(WithRepo(authRepo), WithEncryptor(encryptor), WithJwtGenerator(jwtGenerator))

	jwtKey, err := srv.Login(user)
	if err != nil {
		log.Println(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, key, jwtKey)
}

func TestRegister(t *testing.T){
	authRepo := &mocks.IAuthRepo{}
    encryptor := &mocks.IEncryptor{}

	user := dto.User{
		UserName: "oguz",
		Password: "123wsedrf",
	}
	hashedPass := "123"

	authRepo.On("GetCredentials", user.UserName).Return(model.User{}, nil)
	authRepo.On("AddUser", dto.User{UserName: user.UserName, Password: hashedPass}).Return(nil)
	encryptor.On("Encrypt", user.Password).Return(hashedPass, nil)

	srv := New(WithRepo(authRepo), WithEncryptor(encryptor))
	err := srv.Register(user)
	if err != nil{
		log.Println(err)
	}
	assert.NoError(t, err)
}