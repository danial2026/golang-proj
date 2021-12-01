package services

import (
	"github.com/danial2026/golang-proj/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UsersService = usersService{}
)

type usersService struct{}

func (service usersService) Save(user *domain.User) error {
	if err := user.Save(); err != nil {
		return err
	}
	return nil
}

func (service usersService) GetByEmail(email string) []*domain.User {
	var user domain.User
	users, err := user.GetByEmail(email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return nil
	}
	return users
}
