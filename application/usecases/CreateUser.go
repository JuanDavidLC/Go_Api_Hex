package usecases

import (
	"github.com/JuanDavidLC/Go_Api_Hex/application/commands"
	"github.com/JuanDavidLC/Go_Api_Hex/domain/models"
	"github.com/JuanDavidLC/Go_Api_Hex/domain/ports"
)

type CreateUserInputPort interface {
	CreateUser(userCommand commands.UserCommand) (models.User, error)
}

type CreateUser struct {
	UserRepository ports.UserRepository
}

func NewCreateUser(userRepository ports.UserRepository) *CreateUser {

	return &CreateUser{UserRepository: userRepository}

}

func (c *CreateUser) CreateUser(userCommand commands.UserCommand) (models.User, error) {

	var newUser models.User

	newUser, err := newUser.CreateUser(userCommand.Name, userCommand.Last_name)
	if err != nil {
		return newUser, err
	}

	newUser.Id_User, err = c.UserRepository.Save(&newUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}
