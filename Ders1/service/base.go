package service

import repo "example/repository"

type Service struct{
	repo repo.Repo
}

func New(repo repo.Repo) Service {
	return Service{repo: repo}
}

