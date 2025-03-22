package task

type TaskCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	UserId      int64  `json:"userId" binding:"required"`
}
