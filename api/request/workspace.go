package request

type CreateWorkspace struct {
	Name 		string `json:"name"`
	Description string `json:"description"`
}

type UpdateWorkspace struct {
	Name 		*string `json:"name"`
	Description *string `json:"description"`
}