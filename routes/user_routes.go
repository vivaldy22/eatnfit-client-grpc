package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	auth_service "github.com/vivaldy22/eatnfit-client/proto"
	"github.com/vivaldy22/eatnfit-client/tools/respJson"
	"github.com/vivaldy22/eatnfit-client/tools/vError"
	"github.com/vivaldy22/eatnfit-client/tools/varMux"
)

type userRoute struct {
	service auth_service.UserCRUDClient
}

func NewUserRoute(service auth_service.UserCRUDClient, r *mux.Router) {
	handler := &userRoute{service: service}

	prefix := r.PathPrefix("/users").Subrouter()
	prefix.HandleFunc("", handler.getAll).Methods(http.MethodGet)
	prefix.HandleFunc("", handler.create).Methods(http.MethodPost)
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
	var level *auth_service.User
	err := json.NewDecoder(r.Body).Decode(&level)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		created, err := l.service.Create(context.Background(), level)

		if err != nil {
			vError.WriteError("Create User Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), &auth_service.ID{
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

	data, err := l.service.GetByID(context.Background(), &auth_service.ID{
		Id: id,
	})

	if err != nil {
		vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *userRoute) update(w http.ResponseWriter, r *http.Request) {
	var level *auth_service.User
	err := json.NewDecoder(r.Body).Decode(&level)

	if err != nil {
		vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		id := varMux.GetVarsMux("id", r)

		authID := &auth_service.ID{
			Id: id,
		}

		_, err := l.service.Update(context.Background(), &auth_service.UserUpdateRequest{
			Id:   authID,
			User: level,
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

	authID := &auth_service.ID{
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
