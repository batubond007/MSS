package http

import (
	"errors"
	"net/http"
)

func (s *Server) HandleRegister(w http.ResponseWriter, r *http.Request) {
	user, err := newUser(r)
	if err != nil {
		if errors.Is(err, invalidUser) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = s.sp.IamService().Register(user)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func (s *Server) HandleUnregister(w http.ResponseWriter, r *http.Request) {
	user, err := newUser(r)
	if err != nil {
		if errors.Is(err, invalidUser) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = s.sp.IamService().Unregister(user)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
