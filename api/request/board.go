package request

type CreateBoard struct {
	Name 		string `json:"name"`
	Description string `json:"description"`
	WorkspaceID uint   `json:"workspace_id"`
}

type UpdateBoard struct {
	Name 		*string `json:"name"`
	Description *string `json:"description"`
}