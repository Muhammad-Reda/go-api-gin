package task

type TaskUpdate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	UserId      int64  `json:"user_id" binding:"required"`
}
