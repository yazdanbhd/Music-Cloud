package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yazdanbhd/Music-Cloud/config"
)

type Server struct {
	cfg config.Config
}

func New(cfg config.Config) Server {
	return Server{cfg: cfg}
}

func (s Server) Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/api/users")
	userGroup.POST("/register", s.UserRegister)
	userGroup.POST("/login", s.UserLogin)

	musicGroup := e.Group("/api/music")
	musicGroup.POST("/upload", s.UploadMusic)

	e.Logger.Fatal(e.Start(":8080"))
}
