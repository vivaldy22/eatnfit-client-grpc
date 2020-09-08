package config

import (
	"log"

	authservice "github.com/vivaldy22/eatnfit-client/proto"
	"github.com/vivaldy22/eatnfit-client/tools/viper"
	"google.golang.org/grpc"
)

func newLevelClient() authservice.LevelCRUDClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return authservice.NewLevelCRUDClient(conn)
}

func newTokenClient() authservice.JWTTokenClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return authservice.NewJWTTokenClient(conn)
}
