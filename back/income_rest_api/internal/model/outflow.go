package model

type Outflow struct {
	ID            int `json:"id" db:"id"`
	GoodsID       int `json:"goods_id" db:"goods_id" validate:"required"`
	SectionID     int `json:"section_id" db:"section_id" validate:"required"`
	GoodsCount    int `json:"goods_count" db:"goods_count" validate:"required"`
	ContractorsID int `json:"contractors_id" db:"contractors_id" validate:"required"`
}
