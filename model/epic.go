package model

// Epic represents a card group
type Epic struct {
	BaseModel
	Title string `json:"title" binding:"required"`
}
