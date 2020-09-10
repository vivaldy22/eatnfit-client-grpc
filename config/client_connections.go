package config

import (
	"log"

	"github.com/vivaldy22/eatnfit-client/proto/auth"

	"github.com/vivaldy22/eatnfit-client/tools/viper"
	"google.golang.org/grpc"
)

func newLevelClient() auth.LevelCRUDClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return auth.NewLevelCRUDClient(conn)
}

func newAuthClient() auth.JWTTokenClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return auth.NewJWTTokenClient(conn)
}

func newUserClient() auth.UserCRUDClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return auth.NewUserCRUDClient(conn)
}
