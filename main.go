package main

import (
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
)

func main() {
	dbConfig := mysqldb.Config{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "musiccloudRoo7t0lk2o20",
		DBName:   "music_cloud_db",
	}

	server := httpserver.New(dbConfig)

	server.Run()
}
