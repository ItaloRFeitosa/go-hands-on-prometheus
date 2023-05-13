package main

import (
	"github.com/italorfeitosa/go-hands-on-prometheus/internal"
)

func main() {

	internal.StartServer(internal.NewContainer(), "8080")
}
