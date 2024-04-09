package todo

import (
	"fmt"

	"github.com/thefabric-io/eventsourcing"
)

func haveAnID(aggregate *eventsourcing.Aggregate[*Todo]) error {
	if aggregate.ID() == "" {
		return fmt.Errorf("ID is required")
	}

	return nil
}

func haveValidTasks(aggregate *eventsourcing.Aggregate[*Todo]) error {
	for _, t := range aggregate.State().Tasks {
		if !t.Status.isValid() {
			return fmt.Errorf("task '%s' status is invalid", t.ID)
		}

		if t.Title == "" {
			return fmt.Errorf("task '%s' title is required", t.ID)
		}
	}

	return nil
}
