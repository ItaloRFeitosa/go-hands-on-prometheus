package fiberutils

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var (
		message = "Internal Server Error"
		code    = fiber.StatusInternalServerError
		e       *fiber.Error
	)

	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	} else {
		log.Println("[unknown error] %w", err)
	}

	return ctx.Status(code).JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}
