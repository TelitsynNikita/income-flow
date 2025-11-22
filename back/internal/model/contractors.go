package model

type Contractor struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" validate:"required"`
	INN  string `json:"inn" db:"inn" validate:"required"`
	Type string `json:"type" db:"type" validate:"required"`
}
