package models

import "database/sql"

type Status int

const (
	NotStarted Status = iota
	InProgress
	Completed
	Pending
)

type Task struct {
	ID          string       `json:"id"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Status      string       `json:"status" binding:"required"`
	UserID      string       `json:"userID" binding:"required"`
	CreatedAT   sql.NullTime `json:"createdAt"`
	UpdatedAt   sql.NullTime `json:"updatedAt"`
}

type UpdateTask struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	UserID string `json:"userID"`
}

func (s Status) String() string {
	return [...]string{"Not started", "In progress", "Completed", "Pending"}[s]
}

func (s Status) EnumIndex() int {
	return int(s)
}
