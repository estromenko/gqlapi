package database

import (
	"gqlapi/database/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// UserRepository ...
type UserRepository struct {
	db *dbx.DB
}

// FindAll ...
func (u *UserRepository) FindAll() ([]models.User, error) {
	var userList []models.User

	query := u.db.NewQuery(`SELECT * FROM users`)
	err := query.All(&userList)
	return userList, err
}

// FindByID ...
func (u *UserRepository) FindByID(id string) *models.User {
	var user *models.User

	query := u.db.NewQuery(`SELECT * FROM users WHERE id=` + id)
	query.One(&user)
	return user
}
