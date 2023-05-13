package internal

import (
	"fmt"
	"log"
)

func StartServer(c *Container, port string) {
	if err := c.FiberApp.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
