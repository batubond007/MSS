package main

import (
	"MSS/src/api/http"
)

func main() {
	opt := NewOptions()
	sp := NewServiceProvider(opt)
	server := http.NewServer(sp)
	sp.MessageService().Start()
	err := server.ListenAndServe("localhost:8080")
	if err != nil {
		panic(err)
	}
}
