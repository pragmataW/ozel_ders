package service

import (
	"example/pkg"
	repo "example/repository"
)

type Service struct{
	repo repo.Repo
	encryptor pkg.Encryptor
	jwtGenerator pkg.JwtGenerator
}

type Option func(*Service)

func WithRepo(repo repo.Repo) Option {
	return func(s *Service) {
		s.repo = repo
	}
}

func WithEncryptor(encryptor pkg.Encryptor) Option {
	return func(s *Service) {
		s.encryptor = encryptor
	}
}

func WithJwtGenerator(jwtGenerator pkg.JwtGenerator) Option {
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
