package main

import (
	"fmt"
	"net/http"
)

//обработать все ошибки
//не добавлять сообщение если json плохо распарсился + вывести ошибку

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
