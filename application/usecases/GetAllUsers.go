package usecases

import (
	"github.com/JuanDavidLC/Go_Api_Hex/domain/models"
	"github.com/JuanDavidLC/Go_Api_Hex/domain/ports"
)

type GetAllUsersInputPort interface {
	GetAllUsers() ([]models.User, error)
}

type GetAllUsers struct {
	UserRepository ports.UserRepository
}

func NewGetAllUsers(userRepository ports.UserRepository) *GetAllUsers {

	return &GetAllUsers{UserRepository: userRepository}

}

func (g GetAllUsers) GetAllUsers() ([]models.User, error) {

	users, err := g.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil

}
