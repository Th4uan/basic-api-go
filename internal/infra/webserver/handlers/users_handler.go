package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Th4uan/basic-api-go/internal/dto"
	"github.com/Th4uan/basic-api-go/internal/entity"
	"github.com/Th4uan/basic-api-go/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	UserDB, err := entity.NewUser(user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(UserDB)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
