package iam

import (
	"MSS/src/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service interface {
		Validate(user user.User) (user.User, error)
		Register(user user.User) error
		Unregister(user user.User) error
		List() ([]user.User, error)
	}
	service struct {
		ur user.UserRepository
	}
)

func NewService(ur user.UserRepository) Service {
	return &service{ur: ur}
}

func (s service) Validate(u user.User) (user.User, error) {
	foundUser, err := s.ur.FindByPhone(u.Phone())
	if err != nil {
		return user.User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password()), []byte(u.Password())); err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (s service) Register(u user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password()), 8)
	if err != nil {
		return err
	}
	hashedUser := user.NewUser(u.Phone(), string(hashedPassword))
	if err = s.ur.Save(hashedUser); err != nil {
		return err
	}
	return nil
}

func (s service) Unregister(u user.User) error {
	validUser, err := s.Validate(u)
	if err != nil {
		return nil
	}
	if err = s.ur.Delete(validUser); err != nil {
		return err
	}
	return nil
}

func (s service) List() ([]user.User, error) {
	return s.ur.List()
}
