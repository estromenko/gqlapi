package database

import (
	"database/sql"
	"fmt"
	"gqlapi/models"

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
func (u *UserRepository) FindByID(id string) (*models.User, error) {
	var user *models.User

	query := u.db.NewQuery(`SELECT * FROM users WHERE id=` + id)
	err := query.One(&user)

	if err == sql.ErrNoRows {
		err = nil
	}

	return user, err
}

// Create ...
func (u *UserRepository) Create(user *models.User) error {

	query := u.db.NewQuery(
		fmt.Sprintf(`INSERT INTO users (email, username, password) VALUES ('%s', '%s', '%s') RETURNING id, email, username, password`,
			user.Email, user.Username, user.Password,
		),
	)

	return query.One(&user)
}
