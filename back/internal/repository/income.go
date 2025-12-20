package repository

import (
	"fmt"
	"income-flow/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IncomePostgres struct {
	db *sqlx.DB
}

func NewIncomeRepository(db *sqlx.DB) *IncomePostgres {
	return &IncomePostgres{db: db}
}

func (p *IncomePostgres) CreateGoods(c *fiber.Ctx, good model.Good) (uint, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id uint

	query := fmt.Sprintf("INSERT INTO %s (name, description, volume) VALUES ($1, $2, $3) RETURNING id", goodsTable)
	row := p.db.QueryRowContext(c.Context(), query, good.Name, good.Description, good.Volume)

	if row.Err() != nil {
		return 0, row.Err()
	}

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func (p *IncomePostgres) GetAll(c *fiber.Ctx) ([]model.Good, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var goods []model.Good

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC", goodsTable)
	err = p.db.SelectContext(c.Context(), &goods, query)
	if err != nil {
		return nil, err
	}

	return goods, tx.Commit()
}

func (p *IncomePostgres) Income(c *fiber.Ctx, income model.Income) (uint, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id uint

	query := fmt.Sprintf("INSERT INTO %s (goods_id, section_id, goods_count, contractors_id) VALUES ($1, $2, $3, $4) RETURNING id", incomeTable)
	row := p.db.QueryRowContext(c.Context(), query, income.GoodsID, income.SectionID, income.GoodsCount, income.ContractorsID)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (operation_type, operation_id, goods_count) VALUES ($1, $2, $3)", businessOperationTable)
	row = p.db.QueryRowContext(c.Context(), query, "+", id, income.GoodsCount)
	if row.Err() != nil {
		return 0, row.Err()
	}

	return id, tx.Commit()
}

func (p *IncomePostgres) Outflow(c *fiber.Ctx, outflow model.Outflow) (uint, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id uint

	query := fmt.Sprintf("INSERT INTO %s (goods_id, goods_count, contractors_id, section_id) VALUES ($1, $2, $3, $4) RETURNING id", outflowTable)
	row := p.db.QueryRowContext(c.Context(), query, outflow.GoodsID, outflow.GoodsCount, outflow.ContractorsID, outflow.SectionID)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (operation_type, operation_id, goods_count) VALUES ($1, $2, $3)", businessOperationTable)
	row = p.db.QueryRowContext(c.Context(), query, "-", id, outflow.GoodsCount)
	if row.Err() != nil {
		return 0, row.Err()
	}

	return id, tx.Commit()
}

func (p *IncomePostgres) CreateContractor(c *fiber.Ctx, contractor model.Contractor) (uint, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id uint
	query := fmt.Sprintf("INSERT INTO %s (name, inn, type) VALUES ($1, $2, $3) RETURNING id", contractorTable)
	row := p.db.QueryRowContext(c.Context(), query, contractor.Name, contractor.INN, contractor.Type)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}

func (p *IncomePostgres) GetBusinessOperations(c *fiber.Ctx) ([]model.BusinessOperation, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var allOperations []model.BusinessOperation

	query := fmt.Sprintf("SELECT * FROM %s;", businessOperationTable)
	err = p.db.SelectContext(c.Context(), &allOperations, query)
	if err != nil {
		return nil, err
	}

	var result []model.BusinessOperation
	for _, operation := range allOperations {
		var body struct {
			ID            int    `json:"id" db:"id"`
			GoodsID       int    `json:"goods_id" db:"goods_id"`
			SectionID     int    `json:"section_id" db:"section_id"`
			GoodsCount    int    `json:"goods_count" db:"goods_count"`
			ContractorsID int    `json:"contractors_id" db:"contractors_id"`
			Date          string `json:"date" db:"date"`
		}
		if operation.OperationType == "+" {
			query = fmt.Sprintf("SELECT * FROM %s WHERE id = %d;", incomeTable, operation.OperationID)
			err = p.db.GetContext(c.Context(), &body, query)
			if err != nil {
				return nil, err
			}
		} else if operation.OperationType == "-" {
			query = fmt.Sprintf("SELECT * FROM %s WHERE id = %d;", outflowTable, operation.OperationID)
			err = p.db.GetContext(c.Context(), &body, query)
			if err != nil {
				return nil, err
			}
		}

		operation.GoodsID = body.GoodsID
		operation.SectionID = body.SectionID
		operation.GoodsCount = body.GoodsCount
		operation.ContractorsID = body.ContractorsID

		query = fmt.Sprintf("SELECT * FROM %s WHERE id = %d;", contractorTable, operation.ContractorsID)
		err = p.db.GetContext(c.Context(), &operation.Contractor, query)
		if err != nil {
			return nil, err
		}

		result = append(result, operation)
	}

	return result, tx.Commit()
}

func (p *IncomePostgres) CreateSection(c *fiber.Ctx, section model.Section) (uint, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var id uint

	query := fmt.Sprintf("INSERT INTO %s (volume) VALUES ($1) RETURNING id", sectionsTable)
	row := p.db.QueryRowContext(c.Context(), query, section.Volume)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, tx.Commit()
}
