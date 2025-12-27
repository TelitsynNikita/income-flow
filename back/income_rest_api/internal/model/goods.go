package model

type Good struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required,gt=0,lte=255"`
	Description string `json:"description" db:"description" validate:"required,gt=0,lte=1000"`
	Volume      int    `json:"volume" db:"volume" validate:"required,gt=0"`
}
