package usecases

import (
	"github.com/JuanDavidLC/Go_Api_Hex/domain/models"
	"github.com/JuanDavidLC/Go_Api_Hex/domain/ports"
)

type GetUserByIdInputPort interface {
	GetUserById(id int64) (models.User, error)
}

type GetUserById struct {
	UserRepository ports.UserRepository
}

func NewGetUserById(userRepository ports.UserRepository) *GetUserById {

	return &GetUserById{UserRepository: userRepository}

}

func (g GetUserById) GetUserById(id int64) (models.User, error) {

	user, err := g.UserRepository.GetById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
