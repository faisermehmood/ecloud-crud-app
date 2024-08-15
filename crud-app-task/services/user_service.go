// services/user_service.go
package services

import (
	"crud-app-task/models"
	"crud-app-task/repositories"
	"fmt"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	exists, err := s.userRepository.UserExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("username already exists")
	}
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepository.DeleteUser(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}
