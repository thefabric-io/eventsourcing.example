package main

import "time"

type createTodoRequest struct {
	Name_     string         `json:"name"`
	Metadata_ map[string]any `json:"metadata"`
}

func (c createTodoRequest) Metadata() map[string]any {
	return c.Metadata_
}

func (c createTodoRequest) Name() string {
	return c.Name_
}

type addTodoTaskRequest struct {
	TodoID_      string         `json:"todo_id"`
	Title_       string         `json:"name"`
	Description_ string         `json:"description"`
	DueDate_     *time.Time     `json:"due_date"`
	Metadata_    map[string]any `json:"metadata"`
}

func (c addTodoTaskRequest) Metadata() map[string]any {
	return c.Metadata_
}

func (c addTodoTaskRequest) TodoID() string {
	return c.TodoID_
}

func (c addTodoTaskRequest) Title() string {
	return c.Title_
}

func (c addTodoTaskRequest) Description() string {
	return c.Description_
}

func (c addTodoTaskRequest) DueDate() *time.Time {
	return c.DueDate_
}
