package task

type TaskUpdate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=pending completed"`
	UserId      int64  `json:"userId" binding:"required"`
}
