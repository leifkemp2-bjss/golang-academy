package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() *sql.DB{
	fmt.Println("Creating database")
	// connStr := "postgres://postgres:secret@localhost:5432/tododb?sslmode=disable"
	connStr := "user=academy password=secret host=localhost port=5432 dbname=tododb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not open SQL database")
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Could not ping SQL database")
		log.Fatal(err)
	}

	return db
}

func todosTable(db *sql.DB){
	query := `
	CREATE TABLE IF NOT EXISTS todos(
	id SERIAL PRIMARY KEY
	contents TEXT NOT NULL
	status TEXT NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	println("Todos table created")
}

func insertTodo(db *sql.DB, contents string, status string) int {
	query := `INSERT INTO todos (contents, status) VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(query, contents, status).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}