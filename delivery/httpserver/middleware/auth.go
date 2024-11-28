package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/yazdanbhd/Music-Cloud/delivery/authjwt"
)

func Auth(token authjwt.Token) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return token.VerifyToken(auth)
		},
	})
}
