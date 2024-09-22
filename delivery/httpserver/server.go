package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/api/users")

	userGroup.POST("/register", handler.UserRegister)

	e.Logger.Fatal(e.Start(":8080"))
}
