package database

import (
	"fmt"
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

// Create ...
func (u *UserRepository) Create(email string, username string, password string) (*models.User, error) {
	var user *models.User

	query := u.db.NewQuery(
		fmt.Sprintf(`INSERT INTO users (email, username, password) VALUES ('%s', '%s', '%s') RETURNING id, email, username, password`,
			email, username, password,
		),
	)

	err := query.One(&user)
	if err != nil {
		return nil, err
	}

	return user, err
}
