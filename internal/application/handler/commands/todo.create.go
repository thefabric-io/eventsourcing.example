package commands

import (
	"context"

	"github.com/thefabric-io/eventsourcing"
	"github.com/thefabric-io/eventsourcing.example/internal/application/handler"
	"github.com/thefabric-io/eventsourcing.example/internal/eventstore"
	"github.com/thefabric-io/eventsourcing.example/todo"
	"github.com/thefabric-io/transactional"
)

type CreateTodoCommand interface {
	Name() string
	Metadata() map[string]any
}

type CreateTodoResponse struct {
	Todo *eventsourcing.Aggregate[*todo.Todo]
}

func CreateTodo(transactional transactional.Transactional) handler.Handler[CreateTodoCommand, *CreateTodoResponse] {
	return &createTodo[CreateTodoCommand, *CreateTodoResponse]{
		transactional: transactional,
	}
}

type createTodo[C CreateTodoCommand, R *CreateTodoResponse] struct {
	transactional transactional.Transactional
}

func (h *createTodo[C, R]) Handle(ctx context.Context, cmd C) (R, error) {
	newTodo, err := todo.Create(ctx, &todo.CreateParams{
		Name: cmd.Name(),
	}, cmd.Metadata())
	if err != nil {
		return nil, err
	}

	/*
		Do some other business logic here, for example using other services or other aggregates...
	*/

	tx, err := h.transactional.BeginTransaction(ctx, transactional.DefaultWriteTransactionOptions())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if err := eventstore.Todos().Save(ctx, tx, newTodo); err != nil {
		return nil, err
	}

	/*
		Before committing the transaction, this is an opportune moment to engage
		in additional business logic. Actions such as broadcasting events to a message broker
		or invoking other services could be considered. Utilizing a message broker for event
		publication not only facilitates projections across different services but might also
		be beneficial within the same service for similar purposes. Specifically, this juncture
		allows for the creation of what can be termed 'transactional projections' directly within
		the same database transaction. However, it's important to exercise caution as
		incorporating numerous projections could potentially decelerate the transaction.
		Implementing these in-transaction projections proves advantageous for subsequent validations,
		like ensuring the uniqueness of a todo name by projecting it within the same transaction.
		This aids in preempting duplications in future todo creations by enabling checks against set
		invariants, such as the uniqueness of todo names.
	*/

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &CreateTodoResponse{
		Todo: newTodo,
	}, nil
}
