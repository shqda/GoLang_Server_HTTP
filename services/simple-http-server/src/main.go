package main

import (
	"HttpServer/config"
	"HttpServer/server"
	"log"
)

func main() {
	cfg, err := config.LoadServerConfig()
	if err != nil {
		log.Fatalf("Config loading error: %v", err)
	}
	MyServer := server.MyServer{}
	if err := MyServer.StartServer(cfg); err != nil {
		log.Fatalf("Config loading error: %v", err)
	}
}
