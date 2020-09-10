package config

import (
	"log"
	"net/http"

	"github.com/vivaldy22/eatnfit-client-grpc/middleware"
	"github.com/vivaldy22/eatnfit-client-grpc/routes"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/jwtm"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/viper"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router) {
	port := viper.ViperGetEnv("PORT", "8080")

	log.Printf("Starting Eat N' Fit API Web Server at port: %v\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func InitRouters(r *mux.Router) {
	hmacSampleSecret := viper.ViperGetEnv("HMACSAMPLESECRET", "secret")
	jwtmiddleware := jwtm.NewJWTMiddleware(hmacSampleSecret)
	r.Use(middleware.ActivityLogMiddleware)

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(jwtmiddleware.Handler)

	authClient := newAuthClient()
	userClient := newUserClient()
	levelClient := newLevelClient()
	foodClient := newFoodClient()

	routes.NewAuthRoute(authClient, userClient, r)
	routes.NewLevelRoute(levelClient, admin)
	routes.NewUserRoute(userClient, admin)
	routes.NewFileRoute(r)
	routes.NewFoodRoute(foodClient, admin)
}
