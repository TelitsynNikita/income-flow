package handler

import (
	"fmt"
	"income-flow/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateGoods(c *fiber.Ctx) error {
	var good model.Good

	err := jsoniter.Unmarshal(c.Body(), &good)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to unmarshal json: %v", err))
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&good); err != nil {
		logrus.Error(fmt.Sprintf("failed to validate json: %v", err))
		return fiber.ErrBadRequest
	}

	id, err := h.Service.IncomeService.CreateGoods(c, good)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	goods, err := h.Service.IncomeService.GetAll(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(goods)
}
