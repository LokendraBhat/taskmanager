package services

import (
	"crud-api/internal/models"
	"crud-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(username, password string) (*models.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}