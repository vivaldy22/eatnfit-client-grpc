package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	auth_service "github.com/vivaldy22/eatnfit-client/proto"
	"github.com/vivaldy22/eatnfit-client/tools/respJson"
	"github.com/vivaldy22/eatnfit-client/tools/vError"
	"github.com/vivaldy22/eatnfit-client/tools/varMux"
)

type levelRoute struct {
	service auth_service.LevelCRUDClient
}

func NewLevelRoute(service auth_service.LevelCRUDClient, r *mux.Router) {
	handler := &levelRoute{service: service}

	r.HandleFunc("/levels", handler.getAll).Methods(http.MethodGet)
	prefix := r.PathPrefix("/level").Subrouter()
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
		respJson.WriteJSON(data, w)
	}
}

func (l *levelRoute) create(w http.ResponseWriter, r *http.Request) {
	var level *auth_service.Level
	err := json.NewDecoder(r.Body).Decode(&level)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		created, err := l.service.Create(context.Background(), level)

		if err != nil {
			vError.WriteError("Create Level Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), &auth_service.ID{
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
		data, err := l.service.GetByID(context.Background(), &auth_service.ID{
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
	var level *auth_service.Level
	err := json.NewDecoder(r.Body).Decode(&level)

	if err != nil {
		vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		id := varMux.GetVarsMux("id", r)
		idNum, err := strconv.Atoi(id)
		authID := &auth_service.ID{
			Id: strconv.Itoa(idNum),
		}

		if err != nil {
			vError.WriteError("Converting id failed! not a number", http.StatusExpectationFailed, err, w)
		} else {
			_, err := l.service.Update(context.Background(), &auth_service.LevelUpdateRequest{
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
		authID := &auth_service.ID{
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
