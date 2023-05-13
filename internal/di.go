package internal

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/italorfeitosa/go-hands-on-prometheus/pkg/fiberutils"
	"github.com/italorfeitosa/go-hands-on-prometheus/pkg/gormutils"
)

type Container struct {
	DB *gorm.DB

	FiberApp *fiber.App

	LinkMemoryRepository   *LinkMemoryRepository
	LinkPostgresRepository *LinkPostgresRepository

	LinkController *LinkController

	Base62Codec Base62Codec

	LinkShortenCounter *prometheus.CounterVec

	Metrics *Metrics
}

func NewContainer() *Container {
	c := new(Container)

	provideDB(c)
	provideMetrics(c)
	provideLinkRepository(c)
	provideLinkController(c)
	provideFiberApp(c)

	registerMiddlewares(c)
	registerRoutes(c)

	return c
}

func provideDB(c *Container) {
	c.DB = gormutils.Connect()

	err := c.DB.AutoMigrate(&Link{})
	if err != nil {
		log.Fatal(err)
	}
}

func provideFiberApp(c *Container) {
	c.FiberApp = fiber.New(fiber.Config{
		ErrorHandler: fiberutils.ErrorHandler,
	})
}

func provideMetrics(c *Container) {
	c.Metrics = NewMetrics()
}

func provideLinkRepository(c *Container) {
	c.LinkMemoryRepository = NewMemoryRepository()

	c.LinkPostgresRepository = NewLinkPostgresRepository(c.DB)
}

func provideLinkController(c *Container) {
	c.LinkController = NewLinkController(c.LinkPostgresRepository, c.Base62Codec)
}
