package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(email, password string) (*User, error)
	Login(email, password string) (*User, error)
}

type service struct {
	repo Repostiory
}

func NewService(repo Repostiory) Service {
	return &service{repo: repo}
}

func (s *service) Register(email, password string) (*User, error) {
	_, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, ErrorUserAlreadyExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    email,
		Password: string(hashed),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil

}
func (s *service) Login(email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrorInvalidCredintials
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		return nil, ErrorInvalidCredintials
	}
	user.Password = ""
	return user, nil
}
