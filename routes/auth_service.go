package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	auth_service "github.com/vivaldy22/eatnfit-client/proto"
	"github.com/vivaldy22/eatnfit-client/tools/respJson"
	"github.com/vivaldy22/eatnfit-client/tools/vError"
)

type authService struct {
	service     auth_service.JWTTokenClient
	userService auth_service.UserCRUDClient
}

func NewTokenRoute(service auth_service.JWTTokenClient, userService auth_service.UserCRUDClient, r *mux.Router) {
	handler := &authService{
		service:     service,
		userService: userService,
	}

	prefix := r.PathPrefix("/auth").Subrouter()
	prefix.HandleFunc("/login", handler.login).Methods(http.MethodPost)
	prefix.HandleFunc("/register", handler.register).Methods(http.MethodPost)
}

func (t *authService) login(w http.ResponseWriter, r *http.Request) {
	var user *auth_service.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		data, err := t.userService.GetByEmail(context.Background(), &auth_service.Email{
			Email: user.UserEmail,
		})
		if err != nil {
			vError.WriteError("No User found", http.StatusBadRequest, err, w)
		} else {
			comparePass := bcrypt.CompareHashAndPassword([]byte(data.UserPassword), []byte(user.UserPassword))
			if user.UserEmail == data.UserEmail && comparePass == nil {
				token, err := t.service.GenerateToken(context.Background(), user)

				if err != nil {
					vError.WriteError("Token generation failed!", http.StatusInternalServerError, err, w)
				} else {
					respJson.WriteJSON(&auth_service.LoginResponse{
						User:  data,
						Token: token.Token,
					}, w)
				}
			} else {
				vError.WriteError("Invalid login", http.StatusUnauthorized, err, w)
			}
		}
	}
}

func (t *authService) register(w http.ResponseWriter, r *http.Request) {
	var user *auth_service.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
		user.UserPassword = string(hashedPassword)
		log.Println(user.UserPassword)

		res, err := t.userService.Create(context.Background(), user)

		if err != nil {
			vError.WriteError("Registering failed", http.StatusBadRequest, err, w)
		} else {
			data, err := t.userService.GetByID(context.Background(), &auth_service.ID{Id: res.UserId})

			if err != nil {
				vError.WriteError("Get By ID User failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
