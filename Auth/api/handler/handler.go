package handler

import (
	"auth/service"
)

type Handler struct {
	Auth *service.AuthService
}

func NewHandler(auth *service.AuthService) *Handler {
	return &Handler{
		Auth: auth,	}
}
