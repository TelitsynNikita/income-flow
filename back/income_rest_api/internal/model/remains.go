package model

type Remain struct {
	ID         int `json:"id" db:"id"`
	GoodsID    int `json:"goods_id" db:"goods_id"`
	SectionID  int `json:"section_id" db:"section_id"`
	GoodsCount int `json:"goods_count" db:"goods_count"`
}
