package main

import (
	"net/http"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host		= "localhost"
	port		= 5432
	user		= "a"
	password	= "123"
	dbname		= "db"
)

type User struct {
	ID		int64
	Username	string
	ChosenName	string
	Email		string
	PasswordHash	string
}

func main() {
	// init db
	db, err := init_db()
	if err != nil {
		log.Fatal("connection to db failed:", err)
	}
	defer db.Close()

	// add new user sample
	newUser := &User {
		Username:	"0",
		ChosenName:	"sarsomardo",
		Email:		"fake@fake.com",
		PasswordHash:	"123encrypted",
	}
	err = insert_user(db, newUser)
	if err != nil {
		log.Fatal("insert failed:", err)
	}
	fmt.Println("added user with id", newUser.ID)

	// http stuff
	http.HandleFunc("/health", health)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func init_db() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// check if connection was successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insert_user(db *sql.DB, u *User) error {
	query := `
		INSERT INTO users (username, chosen_name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	err := db.QueryRow(query, u.Username, u.ChosenName, u.Email, u.PasswordHash).Scan(&u.ID)
	return err;
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "up")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you requested %s", r.URL.Path)
}
