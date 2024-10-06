package main

import (
	"context"

	"chipa.me/pkg/handlers"
	"chipa.me/pkg/md"
	"chipa.me/pkg/services"
	"github.com/labstack/gommon/log"
)

func run(ctx context.Context) error {
	file, err := newLogfile(true)
	if err != nil {
		log.Error(err)
	}

	logger := newLogger(file)

	handlers := handlers.NewHandlers(logger)

	postService, err := services.NewPostService()
	if err != nil {
		logger.Error().Err(err).Msg("Init errors in Post Service found")
	}

	md := md.NewMarkdown()

	e := newServer(logger, handlers, postService, md)

	return e.Start(":3000")
}
