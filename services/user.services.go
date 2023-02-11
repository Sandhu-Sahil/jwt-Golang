package services

import "jwt-golang/models"

type UserService interface {
	LoginUser(*models.User) (string, error)
	RegisterUser(*models.User) (string, error)
	GetUserByID(*string) (*models.User, error)
}
