package models

// TIL: tag used for specify the struct fields. It can used for marshaling/unmarshaling
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required" form:"e-mail"`
	Age      int    `json:"age" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
