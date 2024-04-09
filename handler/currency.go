package handler

import (
	"testTaskKMF/schema"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetCurrencyHandler(ctx *fiber.Ctx) error {

	q := new(schema.CurrencyRequest)
	err := ctx.BodyParser(q)

	if err != nil {
		return err
	}
	d, err := h.service.GetCurrencyService(q.Date)

	if err != nil {
		return err
	}

	res := schema.DataResponse{
		Data: d,
	}

	return ctx.XML(res)
}
