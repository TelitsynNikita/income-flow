package model

type BusinessOperation struct {
	ID            int        `json:"id" db:"id"`
	Date          string     `json:"date" db:"date"`
	OperationType string     `json:"operation_type" db:"operation_type"`
	OperationID   int        `json:"operation_id" db:"operation_id"`
	GoodsCount    int        `json:"goods_count" db:"goods_count"`
	GoodsID       int        `json:"goods_id" db:"goods_id"`
	SectionID     int        `json:"section_id" db:"section_id"`
	ContractorsID int        `json:"contractors_id" db:"contractors_id"`
	Contractor    Contractor `json:"contract" db:"contract"`
}
