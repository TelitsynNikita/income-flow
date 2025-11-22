package service

import (
	"income-flow/internal/model"
	"income-flow/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type IncomeService interface {
	CreateGoods(c *fiber.Ctx, good model.Good) (uint, error)
	GetAll(c *fiber.Ctx) ([]model.Good, error)
	CreateContractor(c *fiber.Ctx, contractor model.Contractor) (uint, error)
	Income(c *fiber.Ctx, income model.Income) (uint, error)
	Outflow(c *fiber.Ctx, outflow model.Outflow) (uint, error)
	GetBusinessOperations(c *fiber.Ctx) ([]model.BusinessOperation, error)
	CreateSection(c *fiber.Ctx, section model.Section) (uint, error)
}

type Service struct {
	IncomeService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		IncomeService: NewIncome(repo.IncomeRepository),
	}
}
