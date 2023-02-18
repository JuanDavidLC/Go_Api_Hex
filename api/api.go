package api

import (
	"net/http"

	"github.com/JuanDavidLC/Go_Api_Hex/application/usecases"
	"github.com/JuanDavidLC/Go_Api_Hex/infrastructure/adapters"
	"github.com/JuanDavidLC/Go_Api_Hex/infrastructure/controllers"
	"github.com/JuanDavidLC/Go_Api_Hex/infrastructure/database"
)

func Start() {

	db := database.Connect()
	userPostgresRepository := adapters.NewPostgresRepository(db)
	useCaseCreateUser := usecases.NewCreateUser(userPostgresRepository)
	useCaseGetAllUser := usecases.NewGetAllUsers(userPostgresRepository)
	useCaseGetuserById := usecases.NewGetUserById(userPostgresRepository)
	handler := controllers.NewHandler(useCaseCreateUser, useCaseGetAllUser, useCaseGetuserById)
	initRoutes(handler)

	srv := http.Server{

		Addr: ":8080",
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
