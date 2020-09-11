package routes

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/vivaldy22/eatnfit-client-grpc/tools/respJson"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/vError"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/varMux"

	foodproto "github.com/vivaldy22/eatnfit-client-grpc/proto/food"

	"github.com/vivaldy22/eatnfit-client-grpc/middleware"

	"github.com/gorilla/mux"
)

type packetRoute struct {
	service foodproto.PacketCRUDClient
}

func NewPacketRoute(service foodproto.PacketCRUDClient, r *mux.Router, admin *mux.Router) {
	handler := &packetRoute{service: service}

	adm := admin.PathPrefix("/packets").Subrouter()
	adm.HandleFunc("", handler.getAll).Queries("page", "{page}", "limit", "{limit}", "keyword", "{keyword}").Methods(http.MethodGet)
	adm.HandleFunc("", handler.create).Methods(http.MethodPost)
	adm.HandleFunc("/total", handler.getTotal).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
	adm.HandleFunc("/{id}", handler.update).Methods(http.MethodPut)
	adm.HandleFunc("/{id}", handler.delete).Methods(http.MethodDelete)

	usr := r.PathPrefix("/packets").Subrouter()
	usr.Use(middleware.UsrJwtMiddleware.Handler)
	usr.HandleFunc("", handler.getAll).Queries("page", "{page}", "limit", "{limit}", "keyword", "{keyword}").Methods(http.MethodGet)
	usr.HandleFunc("/total", handler.getTotal).Methods(http.MethodGet)
	usr.HandleFunc("/{id}", handler.getByID).Methods(http.MethodGet)
}

func (l *packetRoute) getAll(w http.ResponseWriter, r *http.Request) {
	var pagination foodproto.Pagination
	pagination.Page = varMux.GetVarsMux("page", r)
	pagination.Limit = varMux.GetVarsMux("limit", r)
	pagination.Keyword = varMux.GetVarsMux("keyword", r)
	data, err := l.service.GetAll(context.Background(), &pagination)

	if err != nil {
		vError.WriteError("Get All Packets Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data.List, w)
	}
}

func (l *packetRoute) getTotal(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetTotal(context.Background(), new(empty.Empty))

	if err != nil {
		vError.WriteError("Get Total Packets Data failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *packetRoute) create(w http.ResponseWriter, r *http.Request) {
	//var packet *foodproto.Packet
	//err := json.NewDecoder(r.Body).Decode(&packet)
	//
	//if err != nil {
	//	vError.WriteError("Decoding json failed!", http.StatusExpectationFailed, err, w)
	//} else {
	//	created, err := l.service.Create(context.Background(), packet)
	//
	//	if err != nil {
	//		vError.WriteError("Create Packet Failed!", http.StatusBadRequest, err, w)
	//	} else {
	//		data, err := l.service.GetByID(context.Background(), &foodproto.ID{
	//			Id: created.PacketId,
	//		})
	//
	//		if err != nil {
	//			vError.WriteError("Get Packet by ID failed", http.StatusBadRequest, err, w)
	//		} else {
	//			respJson.WriteJSON(data, w)
	//		}
	//	}
	//}
	w.Write([]byte("create"))
}

func (l *packetRoute) getByID(w http.ResponseWriter, r *http.Request) {
	id := varMux.GetVarsMux("id", r)
	packetID := &foodproto.ID{
		Id: id,
	}

	data, err := l.service.GetByID(context.Background(), packetID)

	if err != nil {
		vError.WriteError("Get Packet By ID failed!", http.StatusBadRequest, err, w)
	} else {
		respJson.WriteJSON(data, w)
	}
}

func (l *packetRoute) update(w http.ResponseWriter, r *http.Request) {
	//var packet *foodproto.Packet
	//err := json.NewDecoder(r.Body).Decode(&packet)
	//
	//if err != nil {
	//	vError.WriteError("Decoding json failed", http.StatusExpectationFailed, err, w)
	//} else {
	//	id := varMux.GetVarsMux("id", r)
	//	idNum, err := strconv.Atoi(id)
	//	authID := &foodproto.ID{
	//		Id: strconv.Itoa(idNum),
	//	}
	//
	//	if err != nil {
	//		vError.WriteError("Converting id failed! not a number", http.StatusExpectationFailed, err, w)
	//	} else {
	//		_, err := l.service.Update(context.Background(), &foodproto.PacketUpdateRequest{
	//			Id:     authID,
	//			Packet: packet,
	//		})
	//
	//		if err != nil {
	//			vError.WriteError("Updating data failed!", http.StatusBadRequest, err, w)
	//		} else {
	//			data, err := l.service.GetByID(context.Background(), authID)
	//
	//			if err != nil {
	//				vError.WriteError("Get Packet By ID failed!", http.StatusBadRequest, err, w)
	//			} else {
	//				respJson.WriteJSON(data, w)
	//			}
	//		}
	//	}
	//}
	w.Write([]byte("update"))
}

func (l *packetRoute) delete(w http.ResponseWriter, r *http.Request) {
	//id := varMux.GetVarsMux("id", r)
	//idNum, err := strconv.Atoi(id)
	//
	//if err != nil {
	//	vError.WriteError("Converting id failed! not a number", http.StatusExpectationFailed, err, w)
	//} else {
	//	authID := &foodproto.ID{
	//		Id: strconv.Itoa(idNum),
	//	}
	//	data, err := l.service.GetByID(context.Background(), authID)
	//
	//	if err != nil {
	//		vError.WriteError("Get Packet By ID failed!", http.StatusBadRequest, err, w)
	//	} else {
	//		_, err := l.service.Delete(context.Background(), authID)
	//
	//		if err != nil {
	//			vError.WriteError("Delete Packet failed!", http.StatusBadRequest, err, w)
	//		} else {
	//			respJson.WriteJSON(data, w)
	//		}
	//	}
	//}
	w.Write([]byte("delete"))
}
