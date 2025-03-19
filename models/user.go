package models

import "database/sql"

// TIL: tag used for specify the struct fields. It can used for marshaling/unmarshaling
type User struct {
	ID        string       `json:"id"`
	Username  string       `json:"username" binding:"required"`
	Password  string       `json:"password" binding:"required"`
	Email     string       `json:"email" binding:"required" form:"e-mail"`
	CreatedAt sql.NullTime `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Loginuser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
