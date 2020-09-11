package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vivaldy22/eatnfit-client-grpc/middleware"

	"github.com/golang/protobuf/ptypes/empty"

	foodproto "github.com/vivaldy22/eatnfit-client-grpc/proto/food"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/varMux"
)

type foodRoute struct {
	service foodproto.FoodCRUDClient
}

func NewFoodRoute(service foodproto.FoodCRUDClient, r *mux.Router, admin *mux.Router) {
	handler := &foodRoute{service: service}

	adm := admin.PathPrefix("/foods").Subrouter()
	adm.HandleFunc("", handler.getAll).Queries("page", "{page}", "limit", "{limit}", "keyword", "{keyword}").Methods(http.MethodGet)
	adm.HandleFunc("", handler.create).Methods(http.MethodPost)
	adm.HandleFunc("/total", handler.getTotal).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.update).Methods(http.MethodPut)
	adm.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)

	usr := r.PathPrefix("/foods").Subrouter()
	usr.Use(middleware.UsrJwtMiddleware.Handler)
	usr.HandleFunc("", handler.getAll).Queries("page", "{page}", "limit", "{limit}", "keyword", "{keyword}").Methods(http.MethodGet)
	usr.HandleFunc("/total", handler.getTotal).Methods(http.MethodGet)
	usr.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
}

func (l *foodRoute) getAll(w http.ResponseWriter, r *http.Request) {
	var pagination foodproto.Pagination
	pagination.Page = varMux.GetVarsMux("page", r)
	pagination.Limit = varMux.GetVarsMux("limit", r)
	pagination.Keyword = varMux.GetVarsMux("keyword", r)
	data, err := l.service.GetAll(context.Background(), &pagination)

	if err != nil {
		vError.WriteError("Get All Foods Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (l *foodRoute) getTotal(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetTotal(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get Total Foods Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *foodRoute) create(w http.ResponseWriter, r *http.Request) {
	var food *foodproto.Food
	err := json.NewDecoder(r.Body).Decode(&food)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		created, err := l.service.Create(context.Background(), food)

		if err != nil {
			vError.WriteError("Create Food Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), &foodproto.ID{
				Id: created.FoodId,
			})

			if err != nil {
				vError.WriteError("Get Food by ID failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *foodRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := l.service.GetByID(context.Background(), &foodproto.ID{
		Id: id,
	})

	if err != nil {
		vError.WriteError("Get Food By ID failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *foodRoute) update(w http.ResponseWriter, r *http.Request) {
	var food *foodproto.Food
	err := json.NewDecoder(r.Body).Decode(&food)

	if err != nil {
		vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	} else {
		id := varMux.GetVarsMux("id", r)
		authID := &foodproto.ID{
			Id: id,
		}

		_, err := l.service.Update(context.Background(), &foodproto.FoodUpdateRequest{
			Id:   authID,
			Food: food,
		})

		if err != nil {
			vError.WriteError("Updating data failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), authID)

			if err != nil {
				vError.WriteError("Get Food By ID failed!", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *foodRoute) delete(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	authID := &foodproto.ID{
		Id: id,
	}
	data, err := l.service.GetByID(context.Background(), authID)

	if err != nil {
		vError.WriteError("Get Food By ID failed!", http.StatusBadRequest, err, w)
	} else {
		_, err := l.service.Delete(context.Background(), authID)

		if err != nil {
			vError.WriteError("Delete Food failed!", http.StatusBadRequest, err, w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}
