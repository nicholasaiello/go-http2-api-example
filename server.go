package main

import (
	"nicholasaiello/go-middleware-example/routes"
)

const (
	address = ":8080"
)

func main() {
	routes.Run(address)
}
