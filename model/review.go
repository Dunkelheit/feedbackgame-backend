package model

import uuid "github.com/satori/go.uuid"

// Review of a user
type Review struct {
	BaseModel
	UUID       string `json:"uuid"`
	ReviewerID uint   `json:"-"`
	RevieweeID uint   `json:"-"`
	Reviewer   User   `json:"reviewer"`
	Reviewee   User   `json:"reviewee"`
	Cards      []Card `json:"cards" gorm:"many2many:review_cards;"`
	Remark     string `json:"remark"`
	Completed  bool   `json:"completed"`
}

// BeforeCreate callback
func (review *Review) BeforeCreate() (err error) {
	review.UUID = uuid.NewV4().String()
	return
}
