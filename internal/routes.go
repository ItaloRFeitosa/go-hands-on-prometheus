package internal

func registerRoutes(c *Container) {
	c.App.Post("/shorten", c.LinkController.Shorten)
	c.App.Get("/:slug", c.LinkController.Redirect)
}
