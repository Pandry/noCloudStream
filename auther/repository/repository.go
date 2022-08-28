package repository

import "auther/models"

type Repository interface {
	// CreateUser creates the user in the database.
	//	It expects a struct with a password. It will handle the hashing
	CreateUser(models.User) error
	//AuthenticateUser(email, password string) (bool, error)
	GetUserByMail(email string) (models.User, error)
	//FetchUserByID(id string) (models.User, error)
	//UpdateUser(models.User) error
	Close()
}
