package main

import (
	"github.com/joho/godotenv"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"log"
	"os"
	"strconv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConfig := mysqldb.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	server := httpserver.New(dbConfig)

	server.Run()
}
