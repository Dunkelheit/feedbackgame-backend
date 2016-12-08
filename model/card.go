package model

// Card represents a single card
type Card struct {
	BaseModel
	Title    string       `json:"title" binding:"required"`
	Category CardCategory `json:"category" binding:"exists"`
}
