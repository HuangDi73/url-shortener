package auth

import (
	"errors"
	"url-shortener/config"
	"url-shortener/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	*config.Config
	Repo *user.Repository
}

func NewService(repo *user.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) Login(email, password string) (string, error) {
	existedUser, _ := s.Repo.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return existedUser.Email, nil
}

func (s *Service) Register(email, name, password string) (string, error) {
	existedUser, _ := s.Repo.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrNotExists)
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPass),
	}
	createdUser, err := s.Repo.Create(user)
	if err != nil {
		return "", err
	}
	return createdUser.Email, nil
}
