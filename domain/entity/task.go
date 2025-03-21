package entity

import "database/sql"

type Task struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Status      string       `json:"status" binding:"required"`
	UserId      int64        `json:"userId" binding:"required"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   sql.NullTime `json:"updatedAt"`
	DeletedAt   sql.NullTime `json:"deletedAt"`
}
