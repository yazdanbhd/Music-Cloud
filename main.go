package main

import (
	"github.com/yazdanbhd/Music-Cloud/config"
	"github.com/yazdanbhd/Music-Cloud/delivery/httpserver"
)

func main() {

	cfg := config.New("config.yml")
	server := httpserver.New(cfg)

	server.Run()
}
