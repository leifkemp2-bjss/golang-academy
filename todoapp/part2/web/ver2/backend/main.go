package main

// This is version 2 of the web app, it uses only an SQL database and is written to be thread safe

import (
	// "bufio"
	"fmt"
	// "html/template"
	// "log"
	// "net/http"
	"os"
	"os/signal"
	// "strconv"
	// "strings"
	// "flag"
	"context"

	// "academy.com/todoapp/part2/flash"
	// "academy.com/todoapp/todo"
	"academy.com/todoapp/part2/web/database"
)

func main(){
	ctx, ctxDone := context.WithCancel(context.Background())

	database.DB = database.Connect()
	database.CreateTodosTable(database.DB)

	done := startApi(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	ctxDone()
	fmt.Println("user got signal: " + s.String() + " now closing")
	<- done
}