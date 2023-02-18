package ports

import "github.com/JuanDavidLC/Go_Api_Hex/domain/models"

type UserRepository interface {
	Save(user *models.User) (int64, error)
	GetAllUsers() ([]models.User, error)
	GetById(id int64) (models.User, error)
}
