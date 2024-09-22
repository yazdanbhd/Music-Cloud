package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/service/userservice"
	"net/http"
)

func UserRegister(c echo.Context) error {
	var req userservice.RegisterRequest

	fmt.Println("Request user recieved")
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	dbConfig := mysqldb.Config{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "musiccloudRoo7t0lk2o20",
		DBName:   "music_cloud_db",
	}
	db, err := mysqldb.New(dbConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)

	}

	userSvc := userservice.New(db)

	response, err := userSvc.UserRegister(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusCreated, response)
}
