package main

import (
	"github.com/vivaldy22/eatnfit-client-grpc/config"
)

func main() {
	r := config.NewRouter()
	config.InitRouters(r)
	config.RunServer(r)
}
