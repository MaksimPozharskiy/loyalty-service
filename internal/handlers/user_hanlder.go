package handlers

import "github.com/MaksimPozharskiy/loyalty-service/internal/service"

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}
