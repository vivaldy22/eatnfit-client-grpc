package routes

import (
	"context"
	"encoding/json"
	"net/http"

	authproto "github.com/vivaldy22/eatnfit-client-grpc/proto/auth"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/varMux"
)

type userRoute struct {
	service authproto.UserCRUDClient
}

func NewUserRoute(service authproto.UserCRUDClient, r *mux.Router) {
	handler := &userRoute{service: service}

	prefix := r.PathPrefix("/users").Subrouter()
	prefix.HandleFunc("", handler.getAll).Methods(http.MethodGet)
	prefix.HandleFunc("", handler.create).Methods(http.MethodPost)
	prefix.HandleFunc("/email/{email}", handler.getByEmail).Methods(http.MethodGet)
	prefix.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
	prefix.HandleFunc("/{id}", handler.update).Methods(http.MethodPut)
	prefix.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)
}

func (l *userRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetAll(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get All Users Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (l *userRoute) create(w http.ResponseWriter, r *http.Request) {
	var user *authproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
		user.UserPassword = string(hashedPassword)
		created, err := l.service.CreateByAdmin(context.Background(), user)

		if err != nil {
			vError.WriteError("Create User Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), &authproto.ID{
				Id: created.UserId,
			})

			if err != nil {
				vError.WriteError("Get User by ID failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *userRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := l.service.GetByID(context.Background(), &authproto.ID{
		Id: id,
	})

	if err != nil {
		vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *userRoute) getByEmail(w http.ResponseWriter, r *http.Request) {
	email := varMux.GetVarsMux("email", r)

	data, err := l.service.GetByEmail(context.Background(), &authproto.Email{Email: email})

	if err != nil {
		vError.WriteError("Get User By Email failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *userRoute) update(w http.ResponseWriter, r *http.Request) {
	var user *authproto.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		id := varMux.GetVarsMux("id", r)

		authID := &authproto.ID{
			Id: id,
		}

		_, err := l.service.Update(context.Background(), &authproto.UserUpdateRequest{
			Id:   authID,
			User: user,
		})

		if err != nil {
			vError.WriteError("Updating data failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), authID)

			if err != nil {
				vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *userRoute) delete(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	authID := &authproto.ID{
		Id: id,
	}
	data, err := l.service.GetByID(context.Background(), authID)

	if err != nil {
		vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
	} else {
		_, err := l.service.Delete(context.Background(), authID)

		if err != nil {
			vError.WriteError("Delete User failed!", http.StatusBadRequest, err, w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}
