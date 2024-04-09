package todo

import "time"

type task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Status      status     `json:"completion_status"`
}
