package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	mm "github.com/thefueley/scholar-power-api/token"
)

func main() {
	maker, err := mm.NewJWTMaker(os.Getenv("SCHOLAR_POWER_API_SIGNING_KEY"))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Secret Key: %v\n", maker)

	uid := RandomString(10)
	username := "Phobos"
	duration := 10 * time.Minute

	token, payload, err := maker.CreateToken(uid, username, duration)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Token: %v\n", token)
	fmt.Printf("Payload: %v\n", payload)

	payload, err = maker.VerifyToken(token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Payload: %v\n", payload)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
