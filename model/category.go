package model

// CardCategory is the category of the card (positive or negative)
type CardCategory int

const (
	// CardCategoryPositive represents cards with a positive feedback point
	CardCategoryPositive CardCategory = iota

	// CardCategoryNegative represents cards with a negative feedback point
	CardCategoryNegative
)
