package fiberutil

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-hands-on-prometheus/pkg/errs"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var (
		code = errs.HttpStatusCode(err)
	)

	if code == http.StatusInternalServerError {
		log.Printf("[internal error]: %#+v\n", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.Status(code).JSON(err)
}
