package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/vivaldy22/eatnfit-client-grpc/middleware"
	authproto "github.com/vivaldy22/eatnfit-client-grpc/proto/auth"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/varMux"
)

type userRoute struct {
	service authproto.UserCRUDClient
}

func NewUserRoute(service authproto.UserCRUDClient, r *mux.Router, admin *mux.Router) {
	handler := &userRoute{service: service}

	adm := admin.PathPrefix("/users").Subrouter()
	adm.HandleFunc("", handler.getAll).Queries("page", "{page}", "limit", "{limit}", "keyword", "{keyword}").Methods(http.MethodGet)
	adm.HandleFunc("", handler.create).Methods(http.MethodPost)
	adm.HandleFunc("/total", handler.getTotal).Methods(http.MethodGet)
	adm.HandleFunc("/email/{email}", handler.getByEmail).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.update).Methods(http.MethodPut)
	adm.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)

	usr := r.PathPrefix("/users").Subrouter()
	usr.Use(middleware.UsrJwtMiddleware.Handler)
	usr.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)

	topup := r.PathPrefix("/topup").Subrouter()
	topup.Use(middleware.UsrJwtMiddleware.Handler)
	topup.HandleFunc("/{id}", handler.topUp).Methods(http.MethodPost)
	topup.HandleFunc("/history/{id}", handler.topUpHistory).Methods(http.MethodGet)
}

func (l *userRoute) getAll(w http.ResponseWriter, r *http.Request) {
	var pagination = new(authproto.Pagination)
	pagination.Page = varMux.GetVarsMux("page", r)
	pagination.Limit = varMux.GetVarsMux("limit", r)
	pagination.Keyword = varMux.GetVarsMux("keyword", r)
	data, err := l.service.GetAll(context.Background(), pagination)

	if err != nil {
		vError.WriteError("Get All Users Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (l *userRoute) getTotal(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetTotal(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get Total Users Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
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
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
		user.UserPassword = string(hashedPassword)

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

func (l *userRoute) topUp(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	var input *authproto.TopUpInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	} else {
		input.UserId = id
		_, err := l.service.TopUp(context.Background(), input)

		if err != nil {
			vError.WriteError("Top Up Failed!", http.StatusBadRequest, err, w)
		} else {
			data, err := l.service.GetByID(context.Background(), &authproto.ID{
				Id: id,
			})

			if err != nil {
				vError.WriteError("Get User By ID failed!", http.StatusBadRequest, err, w)
			} else {
				respJson.WriteJSON(data, w)
			}
		}
	}
}

func (l *userRoute) topUpHistory(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)

	data, err := l.service.GetBalanceHistory(context.Background(), &authproto.ID{Id: id})

	if err != nil {
		vError.WriteError("Get Balance History Failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}
