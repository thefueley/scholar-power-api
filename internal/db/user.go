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

	insertRow, err := d.ExecContext(ctx,
		`INSERT INTO user (username, password_hash) 
		VALUES 
		($1, $2) RETURNING id`,
		createUserRow.UserName,
		createUserRow.PasswordHash,
	)

	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	if insertRow == nil {
		return "", fmt.Errorf("create user: %w", err)
	}

	return "", nil
}

func (d *Database) GetByID(ctx context.Context, id string) (swoleuser.User, error) {
	var userRow UserRow
	row := d.QueryRowContext(ctx, `SELECT * 
	FROM user
	WHERE id = $1`,
		id,
	)
	err := row.Scan(&userRow.ID, &userRow.UserName, &userRow.PasswordHash)

	if err != nil {
		return swoleuser.User{}, fmt.Errorf("could not get user: %w", err)
	}
	return userRowToUser(userRow), nil
}

func (d *Database) GetByUserName(ctx context.Context, username string) (swoleuser.User, error) {
	var userRow UserRow
	row := d.QueryRowContext(ctx,
		`SELECT id, username, password_hash 
		FROM user 
		WHERE username = $1`, username)
	err := row.Scan(&userRow.ID, &userRow.UserName, &userRow.PasswordHash)

	if err != nil {
		return swoleuser.User{}, fmt.Errorf("could not get user: %w", err)
	}
	return userRowToUser(userRow), nil
}

func (d *Database) UpdateUserPassword(ctx context.Context, uid string, password string) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("update user password: %w", err)
	}

	passwordHash := string(hashedBytes)

	_, err = d.ExecContext(ctx, `UPDATE user SET password_hash = $1 WHERE id = $2`, passwordHash, uid)

	if err != nil {
		return fmt.Errorf("could not update user password: %w", err)
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

func userRowToUser(ur UserRow) swoleuser.User {
	return swoleuser.User{
		ID:           ur.ID.String,
		UserName:     ur.UserName.String,
		PasswordHash: ur.PasswordHash.String,
	}
}
