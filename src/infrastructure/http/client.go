package http

import "net/http"

func NewClient() *http.Client {
	return &http.Client{}
}
