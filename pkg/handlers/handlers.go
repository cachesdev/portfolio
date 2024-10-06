package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Handlers struct {
	// Logger is the only global dependency of handlers.
	// the rest of the deps are passed as arguments
	Logger zerolog.Logger
}

func NewHandlers(logger zerolog.Logger) Handlers {
	return Handlers{
		Logger: logger,
	}
}

func (h Handlers) Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/blog")
	}
}
