package todo

import (
	"time"

	"github.com/thefabric-io/eventsourcing"
)

const TypeTaskAdded = AggregateType + ".task-added.v1" // or "task.added.v1" this semantic is up to you

/*
taskAddedV1 is versioned here, so I can have multiple versions
of the same event, this is really up to you. It becomes easier to
deal with breaking changes for example if something like this happens.
*/
type taskAddedV1 struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

/*
RecordModification is being implemented by taskAddedV1 as taskAddedV1 is modifying
the state of the aggregate.
*/
func (e *taskAddedV1) RecordModification(inceptionTime *time.Time, a *eventsourcing.Aggregate[*Todo]) {
	a.State().UpdatedAt = inceptionTime
}

func (e *taskAddedV1) Type() string {
	return TypeTaskAdded
}

func (e *taskAddedV1) Apply(a *eventsourcing.Aggregate[*Todo]) {
	a.State().AddTask(e.Title, e.Description, e.DueDate, statusPending)
}
