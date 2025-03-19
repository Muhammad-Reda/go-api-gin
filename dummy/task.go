package dummy

import (
	"github.com/muhammad-reda/go-api-gin/models"
)

var Tasks = []models.Task{
	{
		ID:     "1",
		Name:   "Ngoding",
		Status: models.NotStarted.String(),
	},
	{
		ID:     "2",
		Name:   "Ngoding Go",
		Status: models.InProgress.String(),
	},
	{
		ID:     "3",
		Name:   "Ngoding TypeScript",
		Status: models.Completed.String(),
	},
}
