package main

import "HttpServer/src/packages/server"

func main() {
	MyServer := server.MyServer{}
	if err := MyServer.StartServer(); err != nil {
		panic(1)
	}
}
