package models

type Status int

const (
	NotStarted Status = iota
	InProgress
	Completed
	Pending
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type UpdateTask struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (s Status) String() string {
	return [...]string{"Not started", "In progress", "Completed", "Pending"}[s]
}

func (s Status) EnumIndex() int {
	return int(s)
}
