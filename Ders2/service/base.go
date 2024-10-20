package service

import (
	"example/dto"
	"example/model"
)

type IAuthRepo interface {
	GetCredentials(username string) (model.User, error)
	AddUser(dtoUser dto.User) error
}

type IEncryptor interface {
	Decrypt(encryptText string) (string, error)
	Encrypt(plaintext string) (string, error)
}

type IJwtGenerator interface {
	Generate(username string) (string, error)
}

type Service struct {
	repo         IAuthRepo
	encryptor    IEncryptor
	jwtGenerator IJwtGenerator
}

type Option func(*Service)

func WithRepo(repo IAuthRepo) Option {
	return func(s *Service) {
		s.repo = repo
	}
}

func WithEncryptor(encryptor IEncryptor) Option {
	return func(s *Service) {
		s.encryptor = encryptor
	}
}

func WithJwtGenerator(jwtGenerator IJwtGenerator) Option {
	return func(s *Service) {
		s.jwtGenerator = jwtGenerator
	}
}

func New(opts ...Option) Service {
	service := Service{}

	for _, opt := range opts {
		opt(&service)
	}

	return service
}
