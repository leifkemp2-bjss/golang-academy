package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"academy.com/todoapp/todo"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() *sql.DB{
	fmt.Println("Creating database")
	// connStr := "postgres://postgres:secret@localhost:5432/tododb?sslmode=disable"
	connStr := "user=academy password=secret host=localhost port=5432 dbname=tododb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not open SQL database: %s", err)
	}
	// defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping SQL database: %s", err)
	}

	return db
}

func CreateTodosTable(db *sql.DB){
	query := `
	CREATE TABLE IF NOT EXISTS todos(
	id SERIAL PRIMARY KEY,
	contents TEXT NOT NULL,
	status TEXT NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	println("Todos table created")
}

func InsertTodo(db *sql.DB, contents string, status string) int {
	query := `INSERT INTO todos (contents, status) VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(query, contents, status).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func UpdateTodo(db *sql.DB, id int, contents string, status string) {
	var query string
	var err error
	if contents == "" && status == "" {
		log.Fatal("content and status fields have not been provided")
	} else if contents != "" && status == "" {
		query = `UPDATE todos SET contents=($1) WHERE id=($2)`
		_, err = db.Exec(query, contents, id)
	} else if contents == "" && status != "" {
		query = `UPDATE todos SET status=($1) WHERE id=($2)`
		_, err = db.Exec(query, status, id)
	} else {
		query = `UPDATE todos SET contents=($1), status=($2) WHERE id=($3)`
		_, err = db.Exec(query, contents, status, id)
	}

	
	if err != nil {
		log.Fatal(err)
	}
}

func SearchForTodos(db *sql.DB, contents string, status string)(output []todo.Todo) {
	output = []todo.Todo{}

	var query string
	var rows *sql.Rows
	var err error
	if contents == "" && status == "" {
		log.Fatal("content and status fields have not been provided")
	} else if contents != "" && status == "" {
		query = fmt.Sprintf(`SELECT * FROM todos WHERE LOWER(contents) LIKE '%s'`, 
			"%" + strings.ToLower(contents) + "%")
			
		rows, err = db.Query(query)
	} else if contents == "" && status != "" {
		query = `SELECT * FROM todos WHERE status=($1)`
		rows, err = db.Query(query, status)
	} else {
		query = fmt.Sprintf(`SELECT * FROM todos WHERE LOWER(contents) LIKE '%s' AND status=($1)`, 
			"%" + strings.ToLower(contents) + "%")

		rows, err = db.Query(query, status)
	}

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var t todo.Todo
		if err := rows.Scan(&t.Id, &t.Contents, &t.Status); err != nil {
			log.Fatal(err)
		}
		output = append(output, t)
	}

	return output
}

func DeleteTodo(db *sql.DB, id int){
	_, err := db.Exec("DELETE FROM todos WHERE id=($1)", id)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadTodo(db *sql.DB, id int) todo.Todo{
	result := db.QueryRow("SELECT * FROM todos WHERE id=($1)", id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var t todo.Todo
	if err := result.Scan(&t.Id, &t.Contents, &t.Status); err != nil {
		log.Fatal(err)
	}
	return t
}

func ListTodos(db *sql.DB)(output todo.TodoList) {
	output = todo.TodoList{
		List: make(map[int]todo.Todo),
		MaxSize: 100,
	} 
	rows, err := db.Query(`SELECT * FROM todos`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var t todo.Todo
		if err := rows.Scan(&t.Id, &t.Contents, &t.Status); err != nil {
			log.Fatal(err)
		}
		output.List[t.Id] = t
	}

	return output
}