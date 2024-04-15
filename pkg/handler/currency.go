package handler

import (
	"testTaskKMF/pkg/schema"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetCurrencyHandler(ctx *fiber.Ctx) error {

	date := ctx.Query("date")

	d, err := h.service.GetCurrencyService(date)

	if err != nil {
		return err
	}

	res := schema.DataResponse{
		Data: d,
	}

	return ctx.Status(200).JSON(res)
}
