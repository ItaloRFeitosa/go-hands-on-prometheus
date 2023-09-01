package internal

import (
	"strconv"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	requestsTotal *prometheus.CounterVec

	handler fiber.Handler
}

func NewMetrics() *Metrics {
	m := new(Metrics)

	m.requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "conta total de requisições",
		},
		[]string{"code", "method", "path"},
	)

	m.handler = adaptor.HTTPHandler(promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	return m
}

func (m *Metrics) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Next()

		path := c.Route().Path
		if path == "/metrics" {
			return nil
		}

		code := c.Response().StatusCode()
		method := c.Route().Method

		m.requestsTotal.WithLabelValues(strconv.Itoa(code), method, path).Inc()

		return nil
	}
}

func (m *Metrics) Handler() fiber.Handler {
	return m.handler
}
