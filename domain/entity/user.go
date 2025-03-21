package entity

import "database/sql"

type User struct {
	Id        int64        `json:"id"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}
