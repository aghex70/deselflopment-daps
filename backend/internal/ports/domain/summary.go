package domain

type CategorySummary struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	Tasks                uint   `json:"tasks"`
	HighestPriorityTasks uint   `json:"highest_priority_tasks"`
	OwnerID              uint   `json:"owner_id"`
	Shared               uint   `json:"shared"`
}
