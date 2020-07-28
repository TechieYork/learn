package main

import (
	"github.com/sony/gobreaker"
)

func main() {
	gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:"Test",
	})

	return
}
