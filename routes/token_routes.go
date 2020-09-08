package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	auth_service "github.com/vivaldy22/eatnfit-client/proto"
	"github.com/vivaldy22/eatnfit-client/tools/respJson"
	"github.com/vivaldy22/eatnfit-client/tools/vError"
)

type tokenRoute struct {
	service auth_service.JWTTokenClient
}

func NewTokenRoute(service auth_service.JWTTokenClient, r *mux.Router) {
	handler := &tokenRoute{service: service}

	r.HandleFunc("/auth", handler.login).Methods(http.MethodPost)
}

func (t *tokenRoute) login(w http.ResponseWriter, r *http.Request) {
	var user *auth_service.LoginCredentials
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		if user.UserEmail == "admin" && user.UserPassword == "password" {
			token, err := t.service.GenerateToken(context.Background(), user)

			if err != nil {
				vError.WriteError("Token generation failed!", http.StatusInternalServerError, err, w)
			} else {
				respJson.WriteJSON(token, w)
			}
		} else {
			vError.WriteError("Invalid login", http.StatusUnauthorized, err, w)
		}
	}
}
