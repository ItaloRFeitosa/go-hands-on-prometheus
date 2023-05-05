package internal

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Container struct {
	App *fiber.App

	LinkRepository *LinkMemoryRepository

	LinkController *LinkController

	Base62Codec Base62Codec
}

func NewContainer() *Container {
	c := new(Container)

	c.App = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			return ctx.Status(code).JSON(fiber.Map{
				"code":    code,
				"message": message,
			})

		},
	})

	c.App.Use(logger.New())
	c.App.Use(recover.New())

	c.LinkRepository = NewMemoryRepository()

	c.LinkController = NewLinkController(c.LinkRepository, c.Base62Codec)

	registerRoutes(c)

	return c
}

func (c *Container) StartServer(port string) {
	if err := c.App.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
