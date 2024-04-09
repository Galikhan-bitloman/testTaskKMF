package handler

import (
	"testTaskKMF/services"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() (*fiber.App, error) {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	currency := v1.Group("/currency")
	currency.Post("/", CheckDateFormat, h.GetCurrencyHandler)

	return app, nil
}
