package request

type CreateCardGroup struct {
	Name    string `json:"name"`
	BoardID uint   `json:"board_id"`
}

type UpdateCardGroup struct {
	Name *string `json:"name"`
}
