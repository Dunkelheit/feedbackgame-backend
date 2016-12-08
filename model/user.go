package model

// User is somebody like you or me
type User struct {
	BaseModel
	Username   string `json:"username" gorm:"unique_index"`
	FirstName  string `json:"firstName"`
	Surname    string `json:"surname"`
	FullName   string `json:"fullName"`
	JobTitle   string `json:"jobTitle"`
	Department string `json:"department"`
	Company    string `json:"company"`
	Email      string `json:"email" gorm:"unique_index"`
	Avatar     string `json:"avatar"`
	Role       string `json:"role" gorm:"-"`
}
