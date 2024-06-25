package http

import (
	"MSS/src/domain/user"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var invalidUser = errors.New("invalid user")

type userRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func newUser(r *http.Request) (user.User, error) {
	var u userRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return user.User{}, invalidUser
	}
	err = r.Body.Close()
	if err != nil {
		return user.User{}, err
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return user.User{}, err
	}
	return user.NewUser(u.Phone, u.Password), nil
}
