package repository

import (
	"income-flow/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type IncomeRepository interface {
	CreateGoods(c *fiber.Ctx, good model.Good) (uint, error)
	GetAll(c *fiber.Ctx) ([]model.Good, error)
	CreateContractor(c *fiber.Ctx, contractor model.Contractor) (uint, error)
	Income(c *fiber.Ctx, income model.Income) (uint, error)
	Outflow(c *fiber.Ctx, outflow model.Outflow) (uint, error)
	GetBusinessOperations(c *fiber.Ctx) ([]model.BusinessOperation, error)
	CreateSection(c *fiber.Ctx, section model.Section) (uint, error)
}

type Repository struct {
	IncomeRepository
}

func NewRepository(income *sqlx.DB, ) Repository {
	return Repository{
		IncomeRepository: NewIncomeRepository(income),
	}
}
