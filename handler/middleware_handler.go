package handler

import (
	"testTaskKMF/schema"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckDateFormat(ctx *fiber.Ctx) error {

	q := new(schema.CurrencyRequest)

	err := ctx.BodyParser(q)

	if err != nil {
		return err
	}

	_, err = time.Parse("02.02.2006", q.Date)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}
	return ctx.Next()
}
