package main

import (
	"github.com/vivaldy22/eatnfit-client/config"
)

func main() {
	r := config.NewRouter()
	config.RunServer(r)
	config.InitRouters(r)
}
