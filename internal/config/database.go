package config

import (
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		connStr = "host=db user=user password=pass dbname=mydb sslmode=disable"
	}
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func InitDefaultUser() {
	if user := os.Getenv("DEFAULT_USER"); user != "" {
		pass := os.Getenv("DEFAULT_PASS")
		if pass != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err)
			}
			_, err = DB.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2) ON CONFLICT (username) DO NOTHING", user, string(hash))
			if err != nil {
				log.Printf("Failed to insert default user: %v", err)
			}
		}
	}
}