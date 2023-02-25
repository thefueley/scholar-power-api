package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	swoleuser "github.com/thefueley/scholar-power-api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type UserRow struct {
	ID           sql.NullString
	UserName     sql.NullString
	PasswordHash sql.NullString
}

func (d *Database) CreateUser(ctx context.Context, username, password string) (string, error) {
	username = strings.ToLower(username)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	passwordHash := string(hashedBytes)

	createUserRow := UserRow{
		UserName:     sql.NullString{String: username, Valid: true},
		PasswordHash: sql.NullString{String: passwordHash, Valid: true},
	}

	result, err := d.ExecContext(ctx,
		`INSERT INTO user (username, password_hash) 
		VALUES 
		($1, $2) RETURNING id`,
		createUserRow.UserName,
		createUserRow.PasswordHash,
	)

	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	if result == nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	newuserid := d.QueryRowContext(ctx, `SELECT id FROM user WHERE username = $1`, username)
	var uid string
	newuserid.Scan(&uid)

	return uid, nil
}

func (d *Database) GetUserByID(ctx context.Context, id string) (swoleuser.User, error) {
	var userRow UserRow
	row := d.QueryRowContext(ctx, `SELECT id, username 
	FROM user
	WHERE id = $1`,
		id,
	)
	err := row.Scan(&userRow.ID, &userRow.UserName)

	if err != nil {
		return swoleuser.User{}, fmt.Errorf("could not get user: %w", err)
	}
	return userRowToUser(userRow), nil
}

func (d *Database) GetUserByName(ctx context.Context, username string) (swoleuser.User, error) {
	var userRow UserRow
	row := d.QueryRowContext(ctx,
		`SELECT id, username 
		FROM user 
		WHERE username = $1`, username)
	err := row.Scan(&userRow.ID, &userRow.UserName)

	if err != nil {
		return swoleuser.User{}, fmt.Errorf("could not get user: %w", err)
	}
	return userRowToUser(userRow), nil
}

func (d *Database) UpdateUserPassword(ctx context.Context, uid string, password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("model.UpdateUserPassword, GenerateFromPassword: %w", err)
	}

	passwordHash := string(hashedBytes)

	_, err = d.ExecContext(ctx, `UPDATE user SET password_hash = $1 WHERE id = $2`, passwordHash, uid)

	if err != nil {
		fmt.Println("model.UpdateUserPassword, ExecContext: %w", err)
	}

	return nil
}

func (d *Database) DeleteUser(ctx context.Context, id string) error {
	_, err := d.ExecContext(ctx, `DELETE FROM user WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}
	return nil
}

func (d *Database) Login(ctx context.Context, username, password string) (string, error) {
	username = strings.ToLower(username)
	var userRow UserRow
	row := d.QueryRowContext(ctx,
		`SELECT id, username, password_hash 
		FROM user 
		WHERE username = $1`, username)
	err := row.Scan(&userRow.ID, &userRow.UserName, &userRow.PasswordHash)

	if err != nil {
		return "", fmt.Errorf("model.Login, row.Scan: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHash.String), []byte(password))

	if err != nil {
		return "", fmt.Errorf("model.login, CompareHashAndPassword: %w", err)
	}

	return userRow.ID.String, nil
}

func userRowToUser(ur UserRow) swoleuser.User {
	return swoleuser.User{
		ID:           ur.ID.String,
		UserName:     ur.UserName.String,
		PasswordHash: ur.PasswordHash.String,
	}
}
