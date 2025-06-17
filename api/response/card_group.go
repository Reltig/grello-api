package response

import (
	"grello-api/internal/model"
	"grello-api/pkg/collections"
)

type CardGroup struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	BoardID uint   `json:"board_id"`
}

func (resp CardGroup) FromModel(workspace *model.CardGroup) CardGroup {
	return CardGroup{
		ID:      workspace.ID,
		Name:    workspace.Name,
		BoardID: workspace.BoardID,
	}
}

func (resp CardGroup) FromModelCollection(cardGroups []model.CardGroup) []CardGroup {
	return collections.Map(cardGroups, func(cardGroup model.CardGroup) CardGroup { return resp.FromModel(&cardGroup) })
}
