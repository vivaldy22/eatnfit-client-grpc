package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	authproto "github.com/vivaldy22/eatnfit-client-grpc/proto/auth"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/varMux"
)

type levelRoute struct {
	service authproto.LevelCRUDClient
}

func NewLevelRoute(service authproto.LevelCRUDClient, r *mux.Router) {
	handler := &levelRoute{service: service}

	prefix := r.PathPrefix("/levels").Subrouter()
	prefix.HandleFunc("", handler.getAll).Methods(http.MethodGet)
	prefix.HandleFunc("", handler.create).Methods(http.MethodPost)
	prefix.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
	prefix.HandleFunc("/{id}", handler.update).Methods(http.MethodPut)
	prefix.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)
}

func (l *levelRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetAll(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get All Levels Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (l *levelRoute) create(w http.ResponseWriter, r *http.Request) {
	var level *authproto.Level
	err := json.NewDecoder(r.Body).Decode(&level)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		created, err := l.service.Create(context.Background(), level)

		if err != nil {
			vError.WriteError("Create Level Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), &authproto.ID{
				Id: created.LevelId,
			})

			if err != nil {
				vError.WriteError("Get Level by ID failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *levelRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", http.StatusExpectationFailed, err, w)
	} else {
		data, err := l.service.GetByID(context.Background(), &authproto.ID{
			Id: strconv.Itoa(idNum),
		})

		if err != nil {
			vError.WriteError("Get Level By ID failed!", http.StatusBadRequest, err, w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}

func (l *levelRoute) update(w http.ResponseWriter, r *http.Request) {
	var level *authproto.Level
	err := json.NewDecoder(r.Body).Decode(&level)

	if err != nil {
		vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		id := varMux.GetVarsMux("id", r)
		idNum, err := strconv.Atoi(id)
		authID := &authproto.ID{
			Id: strconv.Itoa(idNum),
		}

		if err != nil {
			vError.WriteError("Converting id failed! not a number", http.StatusExpectationFailed, err, w)
		} else {
			_, err := l.service.Update(context.Background(), &authproto.LevelUpdateRequest{
				Id:    authID,
				Level: level,
			})

			if err != nil {
				vError.WriteError("Updating data failed!", http.StatusBadRequest, err, w)
			} else {
				data, err := l.service.GetByID(context.Background(), authID)

				if err != nil {
					vError.WriteError("Get Level By ID failed!", http.StatusBadRequest, err, w)
				} else {
					respJson.WriteJSON(data, w)
				}
			}
		}
	}
}

func (l *levelRoute) delete(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	idNum, err := strconv.Atoi(id)

	if err != nil {
		vError.WriteError("Converting id failed! not a number", http.StatusExpectationFailed, err, w)
	} else {
		authID := &authproto.ID{
			Id: strconv.Itoa(idNum),
		}
		data, err := l.service.GetByID(context.Background(), authID)

		if err != nil {
			vError.WriteError("Get Level By ID failed!", http.StatusBadRequest, err, w)
		} else {
			_, err := l.service.Delete(context.Background(), authID)

			if err != nil {
				vError.WriteError("Delete Level failed!", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}
