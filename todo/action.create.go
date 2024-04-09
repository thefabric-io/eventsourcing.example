package todo

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/thefabric-io/eventsourcing"
)

type CreateParams struct {
	Name string `json:"name"`
}

func Create(ctx context.Context, cmd *CreateParams, metadata map[string]any) (*eventsourcing.Aggregate[*Todo], error) {
	newTodo := eventsourcing.InitZeroAggregate(&Todo{})

	evt := createdV1{
		ID:   "todo_" + ksuid.New().String(),
		Name: cmd.Name,
	}

	eventsourcing.NewEvent[*Todo](&evt, metadata).Apply(newTodo)

	newTodo.Must(haveAnID)

	if err := newTodo.Check(); err != nil {
		return nil, err
	}

	return newTodo, nil
}
