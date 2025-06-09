package response

import (
	"grello-api/internal/model"
	"grello-api/pkg/collections"
)

type Workspace struct {
	ID			uint   `json:"id"`
	Name 		string `json:"name"`
	Description string `json:"description"`
	UserID 		uint   `json:"user_id"`
}

func (resp Workspace) FromModel(workspace *model.Workspace) Workspace {
	return Workspace{
		ID:			 workspace.ID,
		Name:   	 workspace.Name,
		Description: workspace.Description,
		UserID:  	 workspace.UserID,
	}
}

func (resp Workspace) FromModelCollection(workspaces []model.Workspace) []Workspace {
	return collections.Map(workspaces, func (workspace model.Workspace) Workspace { return resp.FromModel(&workspace) })
}
