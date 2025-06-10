package response

import (
	"grello-api/internal/model"
	"grello-api/pkg/collections"
)

type Board struct {
	ID			uint   `json:"id"`
	Name 		string `json:"name"`
	Description string `json:"description"`
	WorkspaceID uint   `json:"workspace_id"`
}

func (resp Board) FromModel(workspace *model.Board) Board {
	return Board{
		ID:			 workspace.ID,
		Name:   	 workspace.Name,
		Description: workspace.Description,
		WorkspaceID: workspace.WorkspaceID,
	}
}

func (resp Board) FromModelCollection(boards []model.Board) []Board {
	return collections.Map(boards, func (board model.Board) Board { return resp.FromModel(&board) })
}
