package routes

import (
	"context"
	"encoding/json"
	"net/http"

	authproto "github.com/vivaldy22/eatnfit-client-grpc/proto/auth"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
)

type authService struct {
	service     authproto.JWTTokenClient
	userService authproto.UserCRUDClient
}

func NewAuthRoute(service authproto.JWTTokenClient, userService authproto.UserCRUDClient, r *mux.Router) {
	handler := &authService{
		service:     service,
		userService: userService,
	}

	prefix := r.PathPrefix("/auth").Subrouter()
	prefix.HandleFunc("/login", handler.login).Methods(http.MethodPost)
	prefix.HandleFunc("/register", handler.register).Methods(http.MethodPost)
}

func (t *authService) login(w http.ResponseWriter, r *http.Request) {
	var user *authproto.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		data, err := t.userService.GetByEmail(context.Background(), &authproto.Email{
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
					respJson.WriteJSON(&authproto.LoginResponse{
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
	var user *authproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
		user.UserPassword = string(hashedPassword)

		res, err := t.userService.Create(context.Background(), user)

		if err != nil {
			vError.WriteError("Registering failed", http.StatusBadRequest, err, w)
		} else {
			data, err := t.userService.GetByID(context.Background(), &authproto.ID{Id: res.UserId})

			if err != nil {
				vError.WriteError("Get By ID User failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
