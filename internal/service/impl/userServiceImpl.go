package impl

import (
	"guitarStore/internal/repository"
	"guitarStore/internal/service"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

type userServiceImpl struct {
	userRepository *repository.UserRepository
}
