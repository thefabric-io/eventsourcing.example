package todo

import (
	"time"

	"github.com/thefabric-io/eventsourcing"
)

const TypeCreated = AggregateType + ".created.v1"

/*
createdV1 is versioned here, so I can have multiple versions
of the same event, this is really up to you. It becomes easier to
deal with breaking changes for example if something like this happens.
*/
type createdV1 struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (e *createdV1) AggregateID() string {
	return e.ID
}

/*
RecordInception is being implemented by createdV1 as createdV1 is an inception event,
meaning that the aggregate is being created with this event.
*/
func (e *createdV1) RecordInception(inceptionTime *time.Time, a *eventsourcing.Aggregate[*Todo]) {
	a.State().CreatedAt = inceptionTime
}

func (e *createdV1) Type() string {
	return TypeCreated
}

func (e *createdV1) Apply(a *eventsourcing.Aggregate[*Todo]) {
	a.State().Tasks = make([]task, 0)
	a.State().Name = e.Name
}
