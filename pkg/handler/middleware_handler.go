package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CheckDateFormat(ctx *fiber.Ctx) error {

	date := ctx.Query("date")

	if date == "" {
		return ctx.Next()
	}

	_, err := time.Parse("02.02.2006", date)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}
	return ctx.Next()
}
