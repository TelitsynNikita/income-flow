package handler

import (
	"fmt"
	"income-flow/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func (h *Handler) ContractorsCreate(c *fiber.Ctx) error {
	var contractor model.Contractor

	err := jsoniter.Unmarshal(c.Body(), &contractor)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to unmarshal json: %v", err))
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&contractor); err != nil {
		logrus.Error(fmt.Sprintf("failed to validate json: %v", err))
		return fiber.ErrBadRequest
	}

	id, err := h.Service.IncomeService.CreateContractor(c, contractor)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
