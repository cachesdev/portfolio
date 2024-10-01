package main

import (
	"context"

	"chipa.me/pkg/handlers"
	"github.com/labstack/gommon/log"
)

func run(ctx context.Context) error {
	file, err := newLogfile(true)
	if err != nil {
		log.Error(err)
	}

	logger := newLogger(file)

	handlers := handlers.NewHandlers(logger)

	e := newServer(logger, handlers)

	return e.Start(":3000")
}
