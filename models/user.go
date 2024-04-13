package models

import (
	"errors"
	"example/rest_api/db"
	"example/rest_api/utils"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

func (e *User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hash, err := utils.HashPassword(e.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Email, hash) //Use .Exec for update, insert, delete data
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)
	
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	
	if err != nil {
		return errors.New("Credentials invalid")
	}
	
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}