package main

import "net/http"

type MyServer struct {
	*http.Server
}

func (s *MyServer) StartServer() {
	srv := &http.Server{
		Addr:    ":5252",
		Handler: NewMyHandler(),
	}
	err := http.ListenAndServe(srv.Addr, srv.Handler)
	if err != nil {
		panic(1)
	}
}
