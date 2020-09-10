// +heroku goVersion go1.14.4
// +heroku install ./cmd/...// +heroku install ./cmd/...
module github.com/vivaldy22/eatnfit-client-grpc

go 1.14

require (
	github.com/auth0/go-jwt-middleware v0.0.0-20200810150920-a32d7af194d1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/mux v1.8.0
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200904185747-39188db58858 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.25.0
)
