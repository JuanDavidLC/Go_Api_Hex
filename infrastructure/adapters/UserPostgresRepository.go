package adapters

import (
	"database/sql"

	"github.com/JuanDavidLC/Go_Api_Hex/domain/models"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *UserPostgresRepository {

	return &UserPostgresRepository{db: db}

}

func (repo UserPostgresRepository) Save(user *models.User) (int64, error) {

	var id int64
	q := `INSERT INTO 
			users (name,last_name)
			VALUES($1,$2)
			RETURNING id_user`

	stmt, err := repo.db.Prepare(q)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	r := stmt.QueryRow(user.Name, user.Last_name)
	err = r.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (repo UserPostgresRepository) GetAllUsers() ([]models.User, error) {

	q := `SELECT id_user,name,last_name
			FROM users`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {

		user := models.User{}

		err := rows.Scan(&user.Id_User, &user.Name, &user.Last_name)

		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil

}

func (repo UserPostgresRepository) GetById(id int64) (models.User, error) {

	user := models.User{}
	q := `SELECT id_user,name,last_name
			FROM users WHERE id_user = $1`

	row := repo.db.QueryRow(q, id)

	err := row.Scan(&user.Id_User, &user.Name, &user.Last_name)
	if err != nil {

		return models.User{}, err
	}

	return user, nil

}
