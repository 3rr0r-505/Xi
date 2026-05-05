package task

import "time"

type Task struct {
	Id          string     `json:"id"`
	ParentId    string     `json:"parent_id"`
	Title       string     `json:"title"`
	Status      Status     `json:"status"`
	CreatedOn   time.Time  `json:"created_on"`
	CompletedOn *time.Time `json:"completed_on,omitempty"`
	Prioritised bool       `json:"prioritised"`
}

type Status string

const (
	StatusPending Status = "pending"
	StatusDone    Status = "done"
)
