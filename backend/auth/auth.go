package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// ===== Типы =====
type UUID string
type Email string
type Username string
type Company string
type PasswordHash string
type Inn string

type User struct {
	ID       UUID         `db:"id" json:"user_id"`
	Email    Email        `db:"email" json:"email" binding:"required"`
	Username Username     `db:"username" json:"username"`
	Company  Company      `db:"company" json:"company"`
	Password PasswordHash `db:"password" json:"password" binding:"required"`
	Inn      Inn          `db:"inn"`
}

type UserToken struct {
	ID       UUID     `json:"user_id"`
	Email    Email    `json:"email"`
	Username Username `json:"username"`
	Company  Company  `json:"company"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("super_secret_key")

// ===== Хеширование пароля =====
func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func checkPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// ===== Регистрация =====
func Register(ctx context.Context, db *sqlx.DB, u *User) error {
	hashed, err := hashPassword(string(u.Password))
	if err != nil {
		return err
	}
	u.Password = PasswordHash(hashed)
	_, err = db.ExecContext(ctx,
		`INSERT INTO users (email, password, username, company, inn) VALUES ($1, $2, $3, $4, $5)`,
		u.Email, u.Password, u.Username, u.Company, u.Inn)
	return err
}

// ===== Аутентификация =====
func Authenticate(ctx context.Context, db *sqlx.DB, realuser *User) (*User, error) {
	var dbuser User

	err := db.GetContext(ctx, &dbuser,
		`SELECT id, email, password, username, company, inn FROM users WHERE email=$1`, realuser.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := checkPassword(string(dbuser.Password), string(realuser.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &dbuser, nil
}

// ===== JWT =====
func GenerateJWT(user *User) (string, error) {
	claims := &UserToken{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Company:  user.Company,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (*UserToken, error) {
	claims := &UserToken{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
