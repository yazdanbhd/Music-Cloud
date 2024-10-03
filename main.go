package main

import (
	"github.com/joho/godotenv"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/repository/s3/minios3"
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

	minioConfig := minios3.Config{
		Endpoint:        "localhost:9000",
		AccessKeyID:     "AXv9sDkbvdPv7uN8TD1e",
		SecretAccessKey: "l0vqlWnBxiQLAcaipuY6lhVeUd81WAQ10LytaJrM",
		UserSSL:         false,
	}

	server := httpserver.New(dbConfig, minioConfig)

	server.Run()
}
