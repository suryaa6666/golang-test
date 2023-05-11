package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {

	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {

	return user, err
}
