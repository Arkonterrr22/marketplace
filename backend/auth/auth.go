package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("super_secret_key")

type User struct {
	ID       int
	Email    string
	Password string
	Userame  string
	Company  string
	Inn      string
}

type Claims struct {
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	Userame string `json:"username"`
	Company string `json:"company"`
	jwt.RegisteredClaims
}

// Регистрация пользователя
func Register(db *sql.DB, email, password, username, company string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Printf("hashed: %s\n", hashed)
	fmt.Printf("email: %s, password: %s username: %s company: %s\n", email, password, username, company)
	res, err := db.Exec("INSERT INTO users (email, password, username, company) VALUES ($1, $2, $3, $4)", email, password, username, company)
	if err != nil {
		log.Println("Insert error:", err)
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected error:", err)
	} else {
		fmt.Println("Rows affected:", rowsAffected)
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var email, password, username, company string
			var createdAt time.Time

			err := rows.Scan(&id, &email, &password, &username, &company, &createdAt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return err
}

func Authenticate(db *sql.DB, email, password string) (*User, error) {
	var user User
	var hashed string

	err := db.QueryRow("SELECT id, email, username, company, password FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.Email, &user.Userame, &user.Company, &hashed)
	if err != nil {
		result := fmt.Sprintf("Invalid credentials: %s", err)
		return nil, errors.New(result)
	}

	fmt.Println(user.Company, user.Email, user.ID, hashed)

	if password != hashed {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// Генерация JWT токена с полной информацией о пользователе
func GenerateJWT(user *User) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:  user.ID,
		Email:   user.Email,
		Userame: user.Userame,
		Company: user.Company,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Парсинг и валидация JWT из токена
func ParseJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// Получение пользователя по ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	var email string
	err := db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Email: email}, nil
}
