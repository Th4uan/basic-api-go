package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Th4uan/basic-api-go/internal/dto"
	"github.com/Th4uan/basic-api-go/internal/entity"
	"github.com/Th4uan/basic-api-go/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JWTExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, JwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          jwt,
		JWTExpiresIn: JwtExpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	UserDB, err := h.UserDB.FindByEmail(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !UserDB.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, err := h.Jwt.Encode(map[string]interface{}{
		"sub": UserDB.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTExpiresIn)).Unix(),
	})

	acessToken := struct {
		AcessToken string `json:"acess_token"`
	}{
		AcessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acessToken)
	w.WriteHeader(http.StatusOK)
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
