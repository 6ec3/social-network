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


func main() {
	// db stuff
	// postgres connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,)

	// open db connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db:", err)
	}
	defer db.Close()

	// verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to db")

	// create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
		);`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("table created or already exists")

	// insert
	var lastInsertId int
	err = db.QueryRow(`INSERT INTO items(name) VALUES ($1) RETURNING id`, "foo").Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted record id:", lastInsertId)

	// read
	var name string
	err = db.QueryRow(`SELECT name FROM items WHERE id=$1`, lastInsertId).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("read record:", name)




	// http stuff
	http.HandleFunc("/health", health)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "up")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you requested %s", r.URL.Path)
}