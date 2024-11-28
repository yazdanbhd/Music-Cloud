package httpserver

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/yazdanbhd/Music-Cloud/config"
	"github.com/yazdanbhd/Music-Cloud/delivery/authjwt"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver/middleware"
)

type Server struct {
	cfg config.Config
}

func New(cfg config.Config) Server {
	return Server{cfg: cfg}
}

func (s Server) Run() {
	e := echo.New()

	// Use echo middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/api/users")
	userGroup.POST("/register", s.UserRegister)
	userGroup.POST("/login", s.UserLogin)

	musicGroup := e.Group("/api/music")
	musicGroup.POST("/upload", s.UploadMusic, middleware.Auth(
		authjwt.New([]byte(`secret-key`), jwt.SigningMethodHS256)))

	e.Logger.Fatal(e.Start(":8080"))
}
