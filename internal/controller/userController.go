package controller

import "guitarStore/internal/service"

type UserController struct {
	service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: *userService}
}
