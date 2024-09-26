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
var tableName string

func Connect() *sql.DB{
	fmt.Println("Creating database")
	// connStr := "postgres://postgres:secret@localhost:5432/tododb?sslmode=disable"
	connStr := "user=academy password=secret host=localhost port=5433 dbname=tododb sslmode=disable"

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

func CreateTodosTable(db *sql.DB, name string){
	tableName = name
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
	id SERIAL PRIMARY KEY,
	contents TEXT NOT NULL,
	status TEXT NOT NULL
	)`, tableName)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s table created\n", tableName)
}

func InsertTodo(db *sql.DB, contents string, status string)(int, error) {
	if contents == "" {
		return -1, fmt.Errorf("contents field has not been provided")
	}
	if status == "" {
		return -1, fmt.Errorf("status field has not been provided")
	}

	query := fmt.Sprintf(`INSERT INTO %s (contents, status) VALUES ($1, $2) RETURNING id`, tableName)

	var id int
	err := db.QueryRow(query, contents, status).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func UpdateTodo(db *sql.DB, id int, contents string, status string)(int, error) {
	var query string
	var err error
	// var t todo.Todo
	var updatedId int
	if contents == "" && status == "" {
		return -1, fmt.Errorf("content and status fields have not been provided")
	} else if contents != "" && status == "" {
		query = fmt.Sprintf(`UPDATE %s SET contents=($1) WHERE id=($2)`, tableName)
		err = db.QueryRow(query, contents, id).Scan(&updatedId)
	} else if contents == "" && status != "" {
		query = fmt.Sprintf(`UPDATE %s SET status=($1) WHERE id=($2)`, tableName)
		err = db.QueryRow(query, status, id).Scan(&updatedId)
	} else {
		query = fmt.Sprintf(`UPDATE %s SET contents=($1), status=($2) WHERE id=($3)`, tableName)
		err = db.QueryRow(query, contents, status, id).Scan(&updatedId)
	}

	
	if err != nil {
		return -1, err
	}
	return updatedId, nil
}

func SearchForTodos(db *sql.DB, contents string, status string)(output []todo.Todo, err error) {
	if contents == "" && status == "" {
		return nil, fmt.Errorf("contents and status have not been provided, please provide at least one")
	}

	output = []todo.Todo{}

	var query string
	var rows *sql.Rows
	if contents == "" && status == "" {
		log.Fatal("content and status fields have not been provided")
	} else if contents != "" && status == "" {
		query = fmt.Sprintf(`SELECT * FROM %s WHERE LOWER(contents) LIKE '%s'`, tableName,
			"%" + strings.ToLower(contents) + "%")
			
		rows, err = db.Query(query)
	} else if contents == "" && status != "" {
		query = fmt.Sprintf(`SELECT * FROM %s WHERE status=($1)`, tableName)
		rows, err = db.Query(query, status)
	} else {
		query = fmt.Sprintf(`SELECT * FROM %s WHERE LOWER(contents) LIKE '%s' AND status=($1)`, tableName,
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

	return output, nil
}

func DeleteTodo(db *sql.DB, id int) (error){
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=($1)", tableName), id)
	if err != nil {
		return err
	}
	return nil
}

func ReadTodo(db *sql.DB, id int) (todo.Todo, error){
	result := db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id=($1)", tableName), id)
	var t todo.Todo
	if err := result.Scan(&t.Id, &t.Contents, &t.Status); err != nil {
		return t, err
	}
	return t, nil
}

func ListTodos(db *sql.DB)(todo.TodoList, error) {
	output := todo.TodoList{
		List: make(map[int]todo.Todo),
		MaxSize: 100,
	} 
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s`, tableName))
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var t todo.Todo
		if err := rows.Scan(&t.Id, &t.Contents, &t.Status); err != nil {
			return output, err
		}
		output.List[t.Id] = t
	}

	return output, nil
}