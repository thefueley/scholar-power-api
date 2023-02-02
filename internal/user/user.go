package user

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/thefueley/scholar-power-api/token"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrNotImplemented = errors.New("not implemented")
)

type User struct {
	ID           string
	UserName     string
	PasswordHash string
}

type UserStore interface {
	CreateUser(context.Context, string, string) (string, error)
	GetByID(context.Context, string) (User, error)
	GetByUserName(context.Context, string) (User, error)
	UpdateUserPassword(context.Context, string, string) error
	DeleteUser(context.Context, string) error
}

type UserService struct {
	Store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{
		Store: store,
	}
}

func (us *UserService) CreateUser(ctx context.Context, username, password string) (string, error) {
	_, err := us.Store.CreateUser(ctx, username, password)
	if err != nil {
		fmt.Printf("error creating user '%v' or user already exists\n", username)
		return "", err
	}
	maker, err := token.NewJWTMaker(os.Getenv("SCHOLAR_POWER_API_SIGNING_KEY"))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	duration := 10 * time.Minute

	token, payload, err := maker.CreateToken(username, duration)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Timestamp: %v: Token: %v\n", time.Now(), token)
	fmt.Printf("Timestamp: %v: Payload: %v\n", time.Now(), payload)
	payload, err = maker.VerifyToken(token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Timestamp: %v: Payload: %v\n", time.Now(), payload)

	return token, nil
}

func (us *UserService) GetByID(ctx context.Context, id string) (User, error) {
	usr, err := us.Store.GetByID(ctx, id)
	if err != nil {
		fmt.Println("error getting user by id", err)
		return User{}, ErrUserNotFound
	}
	return usr, nil
}

func (us *UserService) GetByUserName(ctx context.Context, username string) (User, error) {
	usr, err := us.Store.GetByUserName(ctx, username)
	if err != nil {
		fmt.Println("error getting user by username", err)
		return User{}, ErrUserNotFound
	}
	return usr, nil
}

func (us *UserService) UpdateUserPassword(ctx context.Context, id string, password string) error {
	err := us.Store.UpdateUserPassword(ctx, id, password)
	if err != nil {
		fmt.Println("error updating user password", err)
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(ctx context.Context, id string) error {
	err := us.Store.DeleteUser(ctx, id)
	if err != nil {
		fmt.Println("error deleting user", err)
		return err
	}
	return nil
}
