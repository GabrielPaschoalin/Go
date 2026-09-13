package models

import (
	"errors"
	"gabriel/api/db"
	"gabriel/api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save insere um novo usuário no banco, com a senha já em hash.
func (u *User) Save() error {
	query := `
		INSERT INTO users
		(email, password)
		VALUES
		(?, ?)
	`

	// Preparar a query
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// Gerar hash da senha antes de salvar
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	// Executar a query e guardar o ID gerado
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err
}

// ValidateCredentials confere se o email existe e se a senha informada bate com o hash salvo.
func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users
		WHERE email = ?
	`

	// Buscar usuário pelo email
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	// Comparar a senha informada com o hash salvo
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
