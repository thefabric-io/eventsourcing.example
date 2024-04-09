package application

import (
	"context"
	"os"

	"github.com/thefabric-io/eventsourcing.example/internal/application/handler"
	"github.com/thefabric-io/eventsourcing.example/internal/application/handler/commands"
	"github.com/thefabric-io/eventsourcing/pgeventstore"
	"github.com/thefabric-io/transactional/pgtransactional"
)

func Initialize(ctx context.Context) (*Application, error) {
	if err := pgeventstore.Init(pgeventstore.EventStorageConfig{
		PostgresURL: os.Getenv("TRANSACTIONAL_DATABASE_URL"),
		/*
			Schema is the schema where the events tables will be created. This helps
			keeping the database organized and clear.
		*/
		Schema: "eventstore",
		/*
			Aggregates contains the table name (comma separated) where the events
			will be stored. You also can configure this string in your environment variable
			and introduce it here.
		*/
		Aggregates: "todos",
	}); err != nil {
		return nil, err
	}

	connection, err := pgtransactional.InitSQLXTransactionalConnection(ctx, os.Getenv("TRANSACTIONAL_DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	a := Application{
		Commands: Commands{
			CreateTodo:  commands.CreateTodo(connection),
			AddTodoTask: commands.AddTodoTask(connection),
		},
	}

	return &a, nil
}

// Application is based on the CQRS pattern. It is meant to be used to store all the commands and queries.
type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	//CreateTodo is the CQRS command to create a new todo.
	CreateTodo handler.Handler[commands.CreateTodoCommand, *commands.CreateTodoResponse]

	// AddTodoTask is the CQRS command to add a new task to a todo.
	AddTodoTask handler.Handler[commands.AddTodoTaskCommand, *commands.AddTodoTaskResponse]
}

// Queries is meant to be used to store all the queries. You can add queries here.
type Queries struct{}
