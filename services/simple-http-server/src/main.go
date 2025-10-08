package main

import (
	"HttpServer/config"
	"HttpServer/server"
	"log"
)

func main() {
	cfg, err := config.LoadServerConfig("services/simple-http-server/config/server_config.yaml")
	if err != nil {
		log.Fatalf("Config loading error: %v", err)
	}
	MyServer := server.MyServer{}
	if err := MyServer.StartServer(cfg); err != nil {
		panic(1)
	}
}
