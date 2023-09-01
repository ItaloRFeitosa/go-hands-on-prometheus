package internal

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/italorfeitosa/go-hands-on-prometheus/pkg/fiberutil"
	"github.com/italorfeitosa/go-hands-on-prometheus/pkg/gormutil"
)

type Container struct {
	DB *gorm.DB

	FiberApp *fiber.App

	LinkMemoryDAO   *LinkMemoryDAO
	LinkPostgresDAO *LinkPostgresDAO

	LinkController *LinkHandler

	Base62Codec Base62Codec

	LinkShortenCounter *prometheus.CounterVec

	Metrics *Metrics
}

func NewContainer() *Container {
	c := new(Container)

	provideDB(c)
	provideMetrics(c)
	provideLinkDAO(c)
	provideLinkController(c)
	provideFiberApp(c)

	registerMiddlewares(c)
	registerRoutes(c)

	return c
}

func provideDB(c *Container) {
	c.DB = gormutil.Connect()

	err := c.DB.AutoMigrate(&Link{})
	if err != nil {
		log.Fatal(err)
	}
}

func provideFiberApp(c *Container) {
	c.FiberApp = fiber.New(fiber.Config{
		ErrorHandler: fiberutil.ErrorHandler,
	})
}

func provideMetrics(c *Container) {
	c.Metrics = NewMetrics()
}

func provideLinkDAO(c *Container) {
	c.LinkMemoryDAO = NewLinkMemoryDAO()

	c.LinkPostgresDAO = NewLinkPostgresDAO(c.DB)
}

func provideLinkController(c *Container) {
	c.LinkController = NewLinkHandler(c.LinkPostgresDAO, c.Base62Codec)
}
