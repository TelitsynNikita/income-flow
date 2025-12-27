package handler

import (
	"income-flow/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	router.Use(cors.New())

	goods := router.Group("/goods")
	goods.Post("/create", h.CreateGoods)
	goods.Get("/get_all", h.GetAll)

	operations := router.Group("/operations")
	operations.Post("/income", h.Income)
	operations.Post("/outflow", h.Outflow)
	operations.Get("/get_business_operations", h.GetBusinessOperations)
	operations.Get("/get_remains", h.GetRemains)

	contractors := router.Group("/contractors")
	contractors.Post("/create", h.ContractorsCreate)

	sections := router.Group("/sections")
	sections.Post("/create", h.CreateSection)

	return router
}
