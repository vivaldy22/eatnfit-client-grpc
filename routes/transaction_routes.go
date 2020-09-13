package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client-grpc/middleware"
	foodproto "github.com/vivaldy22/eatnfit-client-grpc/proto/food"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/varMux"
)

type transRoute struct {
	service foodproto.TransactionCRUDClient
}

func NewTransactionRoute(service foodproto.TransactionCRUDClient, r *mux.Router, admin *mux.Router) {
	handler := &transRoute{service: service}

	adm := admin.PathPrefix("/transactions").Subrouter()
	adm.HandleFunc("", handler.getAll).Queries("page", "{page}", "limit", "{limit}", "keyword", "{keyword}").Methods(http.MethodGet)
	adm.HandleFunc("", handler.create).Methods(http.MethodPost)
	adm.HandleFunc("/total", handler.getTotal).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.getByTransID).Methods(http.MethodGet)
	adm.HandleFunc("/users/{id}", handler.getByTransID).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)

	usr := r.PathPrefix("/transactions").Subrouter()
	usr.Use(middleware.UsrJwtMiddleware.Handler)
	usr.HandleFunc("", handler.create).Methods(http.MethodPost)
	usr.HandleFunc("/{id}", handler.getByTransID).Methods(http.MethodGet)
	usr.HandleFunc("/users/{id}", handler.getByTransID).Methods(http.MethodGet)
}

func (l *transRoute) getAll(w http.ResponseWriter, r *http.Request) {
	var pagination = new(foodproto.Pagination)
	pagination.Page = varMux.GetVarsMux("page", r)
	pagination.Limit = varMux.GetVarsMux("limit", r)
	pagination.Keyword = varMux.GetVarsMux("keyword", r)
	data, err := l.service.GetAll(context.Background(), pagination)

	if err != nil {
		vError.WriteError("Get All Transactions Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (l *transRoute) getTotal(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetTotal(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get Total Transactions Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *transRoute) create(w http.ResponseWriter, r *http.Request) {
	var input *foodproto.TransactionInput
	var trans = new(foodproto.Transaction)
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		trans.UserId = input.UserId
		trans.PacketId = input.PacketId
		trans.Portion = input.Portion
		trans.StartDate = input.StartDate
		trans.StartTime = input.StartTime
		trans.Address = input.Address
		trans.PaymentId = input.PaymentId
		created, err := l.service.Create(context.Background(), trans)

		if err != nil {
			vError.WriteError("Create Transaction Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByTransID(context.Background(), &foodproto.ID{
				Id: created.TransId,
			})

			if err != nil {
				vError.WriteError("Get Transaction by ID failed", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *transRoute) getByTransID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := l.service.GetByTransID(context.Background(), &foodproto.ID{
		Id: id,
	})

	if err != nil {
		vError.WriteError("Get Trans By ID failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *transRoute) getByUserID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := l.service.GetByUserID(context.Background(), &foodproto.ID{
		Id: id,
	})

	if err != nil {
		vError.WriteError("Get Trans By User ID failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *transRoute) delete(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	authID := &foodproto.ID{
		Id: id,
	}
	data, err := l.service.GetByTransID(context.Background(), authID)

	if err != nil {
		vError.WriteError("Get Transaction By ID failed!", http.StatusBadRequest, err, w)
	} else {
		_, err := l.service.Delete(context.Background(), authID)

		if err != nil {
			vError.WriteError("Delete Transaction failed!", http.StatusBadRequest, err, w)
		} else {
			respJson.WriteJSON(data, w)
		}
	}
}
