package http

//import user service and message service
import (
	"MSS/src/application"
	"net/http"
)

type Server struct {
	sp *application.ServiceProvider
}

func NewServer(sp *application.ServiceProvider) *Server {
	return &Server{
		sp: sp,
	}
}

func (s *Server) ListenAndServe(port string) error {
	http.HandleFunc("/register", s.HandleRegister)
	http.HandleFunc("/unregister", s.HandleUnregister)
	http.HandleFunc("/sent-list/", s.HandleSentMessages)
	err := http.ListenAndServe(port, nil)
	return err
}
