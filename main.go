package main

import (
	"fmt"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":5252",
		Handler: &MyHandler{},
	}
	err := http.ListenAndServe(srv.Addr, srv.Handler)
	if err != nil {
		fmt.Println(err)
		return
	}
}
