package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/repository/s3/minios3"
)

type Server struct {
	dbConfig    mysqldb.Config
	minioConfig minios3.Config
}

func New(dbConfig mysqldb.Config, minioConfig minios3.Config) Server {
	return Server{dbConfig: dbConfig, minioConfig: minioConfig}
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
