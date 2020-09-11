package middleware

import (
	"github.com/vivaldy22/eatnfit-client-grpc/tools/consts"
	"github.com/vivaldy22/eatnfit-client-grpc/tools/jwtm"
)

var AdmJwtMiddleware = jwtm.NewJWTMiddleware(consts.HMACADM)
var UsrJwtMiddleware = jwtm.NewJWTMiddleware(consts.HMACUSR)
