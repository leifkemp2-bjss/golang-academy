package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"academy.com/todoapp/part2/web/database"
)

func TestMain(m *testing.M){
	ctx, ctxDone := context.WithCancel(context.Background())
	done := createTestServer(ctx)
	
	m.Run()

	ctxDone()
	deleteTestDatabase()
	<- done
}

func createTestServer(ctx context.Context) <- chan struct{}{
	database.DB = database.Connect()
	database.CreateTodosTable(database.DB, "todostest")
	seedTestDatabase()

	done := startApi(ctx)
	return done
}

func seedTestDatabase(){
	fmt.Println("Seeding DB")
	database.DB.Exec("TRUNCATE TABLE todostest RESTART IDENTITY")

	seedData := []struct{
		contents string
		status string
	}{
		{"Test Todo 1", "To Do"},
		{"Test Todo 2", "In Progress"},
		{"Test Todo 3", "Completed"},
	}

	for _, todo := range seedData{
		_, err := database.InsertTodo(database.DB, todo.contents, todo.status)
		if err != nil {
			log.Fatal(err)
		}
	}
	
}

func deleteTestDatabase(){
	database.DB.Exec("DROP TABLE todostest")
}