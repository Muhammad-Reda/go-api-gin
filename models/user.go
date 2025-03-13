package models

// TIL: tag used for specify the struct fields. It can used for marshaling/unmarshaling
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
