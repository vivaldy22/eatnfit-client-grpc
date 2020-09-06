package routes

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	auth_service "github.com/vivaldy22/eatnfit-client/proto"
	"github.com/vivaldy22/eatnfit-client/tools/respJson"
	"github.com/vivaldy22/eatnfit-client/tools/vError"
	"net/http"
)

type levelRoute struct {
	service auth_service.LevelCRUDClient
}

func NewLevelRoute(service auth_service.LevelCRUDClient, r *mux.Router) {
	handler := &levelRoute{service: service}

	r.HandleFunc("/levels", handler.getAll).Methods(http.MethodGet)
}

func (l *levelRoute) getAll(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetAll(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get All Levels Data failed!", err, &w)
	} else {
		respJson.WriteJSON(data, w)
	}
}
