package models

import (
	"errors"

	"github.com/JuanDavidLC/Go_Api_Hex/domain/validators"
)

type User struct {
	Id_User   int64
	Name      string
	Last_name string
}

func (user *User) CreateUser(name string, last_name string) (User, error) {

	if validators.IsFieldEmpty(name) {

		return User{}, errors.New("name should have some value")

	}

	if validators.IsFieldEmpty(last_name) {

		return User{}, errors.New("LastName has not been sent")

	}

	return User{
		Name:      name,
		Last_name: last_name,
	}, nil

}
