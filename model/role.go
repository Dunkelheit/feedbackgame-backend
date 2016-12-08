package model

// Role of a username
type Role struct {
	BaseModel
	Username string `json:"username" gorm:"unique_index"`
	Role     string `json:"role"`
}
