package main

import (
	"HttpServer/server"
)

func main() {
	MyServer := server.MyServer{}
	if err := MyServer.StartServer(); err != nil {
		panic(1)
	}
}
