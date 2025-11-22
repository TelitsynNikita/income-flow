package model

type Section struct {
	ID     int `json:"id" db:"id"`
	Volume int `json:"volume" db:"volume" validate:"required"`
}
