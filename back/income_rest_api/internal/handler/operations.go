package handler

import (
	"fmt"
	"income-flow/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Income(c *fiber.Ctx) error {
	var income model.Income

	err := jsoniter.Unmarshal(c.Body(), &income)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to unmarshal json: %v", err))
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&income); err != nil {
		logrus.Error(fmt.Sprintf("failed to validate json: %v", err))
		return fiber.ErrBadRequest
	}

	id, err := h.Service.IncomeService.Income(c, income)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to income id: %v", err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handler) Outflow(c *fiber.Ctx) error {
	var outflow model.Outflow

	err := jsoniter.Unmarshal(c.Body(), &outflow)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to unmarshal json: %v", err))
		return fiber.ErrBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err = validate.Struct(&outflow); err != nil {
		logrus.Error(fmt.Sprintf("failed to validate json: %v", err))
		return fiber.ErrBadRequest
	}

	id, err := h.Service.IncomeService.Outflow(c, outflow)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to outflow id: %v", err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handler) GetBusinessOperations(c *fiber.Ctx) error {
	businessOperations, err := h.Service.IncomeService.GetBusinessOperations(c)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to get bussines operations: %v", err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(businessOperations)
}

func (h *Handler) GetRemains(c *fiber.Ctx) error {
	remains, err := h.Service.IncomeService.GetRemains(c)
	if err != nil {
		logrus.Error(fmt.Sprintf("failed to get remains: %v", err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(remains)
}
