package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func (s *Server) HandleSentMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		pageSize = 10
	}
	user, err := newUser(r)
	if err != nil {
		if errors.Is(err, invalidUser) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	res, err := s.sp.MessageService().ListSentMessages(user, pageSize)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	response := NewMessageListResponse(res)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
