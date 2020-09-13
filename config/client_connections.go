package config

import (
	"log"

	foodproto "github.com/vivaldy22/eatnfit-client-grpc/proto/food"

	authproto "github.com/vivaldy22/eatnfit-client-grpc/proto/auth"

	"github.com/vivaldy22/eatnfit-client-grpc/tools/viper"
	"google.golang.org/grpc"
)

func newLevelClient() authproto.LevelCRUDClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return authproto.NewLevelCRUDClient(conn)
}

func newAuthClient() authproto.JWTTokenClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return authproto.NewJWTTokenClient(conn)
}

func newUserClient() authproto.UserCRUDClient {
	host := viper.ViperGetEnv("GRPC_AUTH_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_AUTH_PORT", "1010")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return authproto.NewUserCRUDClient(conn)
}

func newFoodClient() foodproto.FoodCRUDClient {
	host := viper.ViperGetEnv("GRPC_FOOD_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_FOOD_PORT", "1011")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return foodproto.NewFoodCRUDClient(conn)
}

func newPacketClient() foodproto.PacketCRUDClient {
	host := viper.ViperGetEnv("GRPC_FOOD_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_FOOD_PORT", "1011")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return foodproto.NewPacketCRUDClient(conn)
}

func newTransactionClient() foodproto.TransactionCRUDClient {
	host := viper.ViperGetEnv("GRPC_FOOD_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_FOOD_PORT", "1011")
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return foodproto.NewTransactionCRUDClient(conn)
}
