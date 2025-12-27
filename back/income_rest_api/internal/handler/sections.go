package handler

import (
	"fmt"
	"income-flow/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateSection(c *fiber.Ctx) error {
	var section model.Section

	err := jsoniter.Unmarshal(c.Body(), &section)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to unmarshal json: %v", err))
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&section); err != nil {
		logrus.Error(fmt.Sprintf("failed to validate json: %v", err))
		return fiber.ErrBadRequest
	}

	id, err := h.Service.IncomeService.CreateSection(c, section)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
