package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/service/userservice"
	"net/http"
)

func UserRegister(c echo.Context) error {
	var req userservice.RegisterRequest

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

func UserLogin(c echo.Context) error {
	var req userservice.LoginRequest

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

	response, err := userSvc.UserLogin(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, response)
}
