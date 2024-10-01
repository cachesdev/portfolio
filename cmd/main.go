package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}
