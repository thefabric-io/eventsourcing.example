package todo

import (
	"context"
	"time"

	"github.com/thefabric-io/eventsourcing"
)

type AddTaskParams struct {
	Title       string     `json:"name"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
}

func AddTask(ctx context.Context, p *AddTaskParams, existingTodo *eventsourcing.Aggregate[*Todo], metadata map[string]any) error {
	evt := taskAddedV1{
		Title:       p.Title,
		Description: p.Description,
		DueDate:     p.DueDate,
	}

	eventsourcing.NewEvent[*Todo](&evt, metadata).Apply(existingTodo)

	existingTodo.Must(haveValidTasks)

	if err := existingTodo.Check(); err != nil {
		return err
	}

	return nil
}
