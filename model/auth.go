package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jihanlugas/goproject.git/config"
	"log"
	"time"
)

type Auth struct {
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
}

type Credentials struct {
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
}

type Claims struct {
	Email	string	`json:"email"`
	jwt.StandardClaims
}

func (c *Credentials) Signin() error {
	log.Println("Model Signin")
	db := config.DbConn()
	defer db.Close()

	var u User

	err := db.QueryRow("SELECT id, email, password, name FROM users where email = ? AND password = ?",
		c.Email, c.Password).Scan(&u.ID, &u.Email, &u.Password, &u.Name)

	return err
}

func (c *Credentials) GenerateToken() (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Email:  c.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SigningString()
}