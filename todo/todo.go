package todo

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/thefabric-io/eventsourcing"
)

const AggregateType = "todo"

type Todo struct {
	Name      string     `json:"name"`
	Tasks     []task     `json:"tasks"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

/*
StorageName returns the name of the table or collection where the events are stored.
If Todo do not implement this method, the default name used will be the AggregateType.
By convention in some database server like Postgres, the table name should be in plural.
This is why implementing this method is made possible here to override the table name.
*/
func (t *Todo) StorageName() string {
	return "todos"
}

func (t *Todo) Type() string {
	return AggregateType
}

func (t *Todo) Zero() eventsourcing.AggregateState {
	return &Todo{}
}

func (t *Todo) AddTask(title, description string, dueDate *time.Time, status status) {
	if t.Tasks == nil {
		t.Tasks = make([]task, 0)
	}

	newTask := task{
		ID:          fmt.Sprintf("tsk_%s", ksuid.New()),
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}

	t.Tasks = append(t.Tasks, newTask)
}
