package eventstore

import (
	"github.com/thefabric-io/eventsourcing"
	"github.com/thefabric-io/eventsourcing.example/todo"
	"github.com/thefabric-io/eventsourcing/pgeventstore"
)

func Todos() eventsourcing.EventStore[*todo.Todo] {
	return pgeventstore.Storage[*todo.Todo]()
}
