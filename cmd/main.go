package main

import (
	"MSS/src/api/http"
	"MSS/src/application"
)

func main() {
	opt := application.NewOptions()
	sp := application.NewServiceProvider(opt)
	server := http.NewServer(sp)
	sp.MessageService().Start()
	err := server.ListenAndServe("localhost:8080")
	if err != nil {
		panic(err)
	}
}
