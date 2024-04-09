package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/thefabric-io/eventsourcing.example/internal/application"
)

func main() {
	// Load the environment variables.
	_ = godotenv.Load()

	// Create a context for passing across boundaries.
	ctx := context.TODO()

	/*
		Initialize the CQRS application. This line will also initialize
		the event store and create the necessary tables.
	*/
	app, err := application.Initialize(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/*
		let's say this is from an HTTP request body that you unmarshalled
		into the struct createTodoRequest that implements commands.CreateTodoCommand
		interface.
	*/
	createRequest := &createTodoRequest{
		Name_: "My first todo",
		Metadata_: map[string]any{
			"creator":  "user_id_for_example",
			"token":    "jwt_access_token_for_example",
			"trace_id": "trace_2eoSfNKBycphtF71lGXxLAsqrm5",
		},
	}

	// Calling the CreateTodo command to create a new todo.
	createTodoResult, err := app.Commands.CreateTodo.Handle(ctx, createRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n1. First todo with id '%s' created successfully.\n", createTodoResult.Todo.ID())

	/*
		A second HTTP request comes to add a task to a todo. The request body is unmarshalled.
	*/
	afterTomorrow := time.Now().Add(24 * time.Hour)
	addTaskRequest := addTodoTaskRequest{
		// The previous todo ID is being used directly here for this example.
		// In reality the ID is passed through the HTTP request.
		TodoID_:      createTodoResult.Todo.ID(),
		Title_:       "My first task",
		Description_: "This is my first task",
		DueDate_:     &afterTomorrow,
		Metadata_:    nil, // the metadata can also be nil
	}

	// Calling the AddTodoTask command to add a task to the first todo.
	addTodoTaskResponse, err := app.Commands.AddTodoTask.Handle(ctx, addTaskRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("2. Task: \n\t=>'%s'\n added successfully to the first todo item.\n", addTodoTaskResponse.Todo.State().Tasks)

	// Serialize the todo item to JSON and print it.
	b, err := json.Marshal(addTodoTaskResponse.Todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nSerialized todo item (latest state): \n%s\n", string(b))
}
