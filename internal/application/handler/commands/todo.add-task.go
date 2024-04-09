package commands

import (
	"context"
	"time"

	"github.com/thefabric-io/eventsourcing"
	"github.com/thefabric-io/eventsourcing.example/internal/application/handler"
	"github.com/thefabric-io/eventsourcing.example/internal/eventstore"
	"github.com/thefabric-io/eventsourcing.example/todo"
	"github.com/thefabric-io/transactional"
)

/*
AddTodoTaskCommand interface can be replaced by a struct instead.
I like to make use of interface because in some system an HTTP API,
a GRPC API, or a command line API could be used to send commands,
you just have to implement the interface.
*/
type AddTodoTaskCommand interface {
	TodoID() string
	Title() string
	Description() string
	DueDate() *time.Time
	Metadata() map[string]any
}

type AddTodoTaskResponse struct {
	Todo *eventsourcing.Aggregate[*todo.Todo]
}

func AddTodoTask(transactional transactional.Transactional) handler.Handler[AddTodoTaskCommand, *AddTodoTaskResponse] {
	return &addTodoTask[AddTodoTaskCommand, *AddTodoTaskResponse]{
		transactional: transactional,
	}
}

type addTodoTask[C AddTodoTaskCommand, R *AddTodoTaskResponse] struct {
	transactional transactional.Transactional
}

func (h *addTodoTask[C, R]) Handle(ctx context.Context, cmd C) (R, error) {
	tx, err := h.transactional.BeginTransaction(ctx, transactional.DefaultWriteTransactionOptions())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Loads the last versioned todo from the event store
	existingTodo, err := eventstore.Todos().Load(ctx, tx, cmd.TodoID(), eventsourcing.LastVersion)
	if err != nil {
		return nil, err
	}

	if err := todo.AddTask(ctx,
		&todo.AddTaskParams{
			Title:       cmd.Title(),
			Description: cmd.Description(),
			DueDate:     cmd.DueDate(),
		},
		existingTodo,
		cmd.Metadata(),
	); err != nil {
		return nil, err
	}

	if err := eventstore.Todos().Save(ctx, tx, existingTodo); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &AddTodoTaskResponse{
		Todo: existingTodo,
	}, nil
}
