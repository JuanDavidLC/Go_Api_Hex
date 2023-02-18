package database

import (
	"database/sql"
	"fmt"

	"github.com/JuanDavidLC/Go_Api_Hex/envs"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	host := envs.Get("DB_HOST")
	port := envs.Get("DB_PORT")
	database := envs.Get("DB_DATABASE")
	username := envs.Get("DB_USERNAME")
	password := envs.Get("DB_PASSWORD")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db

}
