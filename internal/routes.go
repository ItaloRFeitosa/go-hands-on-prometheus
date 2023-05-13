package internal

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerRoutes(c *Container) {
	c.FiberApp.Get("/metrics", c.Metrics.Handler())
	c.FiberApp.Post("/shorten", c.LinkController.Shorten)
	c.FiberApp.Get("/:slug", c.LinkController.Redirect)
}

func registerMiddlewares(c *Container) {
	c.FiberApp.Use(c.Metrics.Middleware())
	c.FiberApp.Use(logger.New())
	c.FiberApp.Use(recover.New())
}
