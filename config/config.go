package config

import (
	"log"
	"net/http"

	"github.com/vivaldy22/eatnfit-client/tools/jwtm"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/eatnfit-client/middleware"
	"github.com/vivaldy22/eatnfit-client/routes"
	"github.com/vivaldy22/eatnfit-client/tools/viper"
)

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router) {
	host := viper.ViperGetEnv("API_HOST", "localhost")
	port := viper.ViperGetEnv("API_PORT", "8080")

	log.Printf("Starting Eat N' Fit API Web Server at %v port: %v\n", host, port)
	if err := http.ListenAndServe(host+":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func InitRouters(r *mux.Router) {
	hmacSampleSecret := viper.ViperGetEnv("HMACSAMPLESECRET", "secret")
	jwtmiddleware := jwtm.NewJWTMiddleware(hmacSampleSecret)
	r.Use(middleware.ActivityLogMiddleware)

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(jwtmiddleware.Handler)

	tokenClient := newTokenClient()
	userClient := newUserClient()
	levelClient := newLevelClient()

	routes.NewTokenRoute(tokenClient, userClient, r)
	routes.NewLevelRoute(levelClient, admin)
	routes.NewUserRoute(userClient, admin)
}
