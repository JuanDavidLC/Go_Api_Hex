package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanDavidLC/Go_Api_Hex/application/commands"
	"github.com/JuanDavidLC/Go_Api_Hex/application/usecases"
)

type UserHandler interface {
	Store(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request, id int64)
}

type Handler struct {
	CreateUserInputPort  usecases.CreateUserInputPort
	GetAllUsersInputPort usecases.GetAllUsersInputPort
	GetUserByIdInputPort usecases.GetUserByIdInputPort
}

func NewHandler(u usecases.CreateUserInputPort, getall usecases.GetAllUsersInputPort, get usecases.GetUserByIdInputPort) *Handler {

	return &Handler{CreateUserInputPort: u, GetAllUsersInputPort: getall, GetUserByIdInputPort: get}

}

func (h Handler) Store(w http.ResponseWriter, r *http.Request) {

	var UserCommand commands.UserCommand

	err := json.NewDecoder(r.Body).Decode(&UserCommand)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return

	}

	newUser, err := h.CreateUserInputPort.CreateUser(UserCommand)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {

	users, err := h.GetAllUsersInputPort.GetAllUsers()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func (h Handler) GetById(w http.ResponseWriter, r *http.Request, id int64) {

	user, err := h.GetUserByIdInputPort.GetUserById(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
