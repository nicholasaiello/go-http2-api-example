package main

import "github.com/nicholasaiello/go-http2-api-example/routes"

const (
	address = ":8080"
)

func main() {
	routes.Run(address)
}
