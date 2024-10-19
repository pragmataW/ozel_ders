package controller

import "example/service"

type Controller struct {
	service service.Service 
}

func New(service service.Service) Controller {
	return Controller{
		service: service,
	}
}

type loginReq struct{
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}