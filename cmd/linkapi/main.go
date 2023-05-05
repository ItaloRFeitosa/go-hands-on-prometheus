package main

import "github.com/italorfeitosa/go-hands-on-prometheus/internal"

func main() {
	c := internal.NewContainer()

	c.StartServer("8080")
}
