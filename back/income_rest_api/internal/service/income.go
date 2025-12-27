package service

import (
	"income-flow/internal/model"
	"income-flow/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type Income struct {
	repo repository.IncomeRepository
}

func NewIncome(repo repository.IncomeRepository) *Income {
	return &Income{repo}
}

func (i *Income) CreateGoods(c *fiber.Ctx, good model.Good) (uint, error) {
	return i.repo.CreateGoods(c, good)
}

func (i *Income) GetAll(c *fiber.Ctx) ([]model.Good, error) {
	return i.repo.GetAll(c)
}

func (i *Income) CreateContractor(c *fiber.Ctx, contractor model.Contractor) (uint, error) {
	return i.repo.CreateContractor(c, contractor)
}

func (i *Income) Income(c *fiber.Ctx, income model.Income) (uint, error) {
	return i.repo.Income(c, income)
}

func (i *Income) Outflow(c *fiber.Ctx, outflow model.Outflow) (uint, error) {
	return i.repo.Outflow(c, outflow)
}

func (i *Income) GetBusinessOperations(c *fiber.Ctx) ([]model.BusinessOperation, error) {
	return i.repo.GetBusinessOperations(c)
}

func (i *Income) CreateSection(c *fiber.Ctx, section model.Section) (uint, error) {
	return i.repo.CreateSection(c, section)
}

func (i *Income) GetRemains(c *fiber.Ctx) ([]model.Remain, error) {
	return i.repo.GetRemains(c)
}
