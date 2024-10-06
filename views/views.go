package views

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	err := t.Render(ctx.Request().Context(), buf)
	if err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func CurrentYear() string {
	return strconv.Itoa(time.Now().Year())
}

func GeneratePhrase() string {
	phrases := []string{
		"Giving Unwated Advice Since 1999",
		"This Blog Kinda Smells Like Cigarettes...",
		"Caution: May Contain Traces of Javascript",
		"Loading... Please Be Patient",
		"Under Construction Since Forever",
		"Certified Y2K Compliant",
		"Rad Ideas, Totally Tubular Design",
		"This Space Intentionally Left Blank",
		"Now With 30% More Pixels!",
		"As Seen on MTV",
		"Batteries Not Included",
		"Please Rewind Before Returning",
		"Don't Lick the Screen (We're Watching You)",
		"If You Can Read This, You're Procrastinating",
		"Guaranteed to Work 60% of the Time, Every Time",
		"Load Times Faster Than Your Dad's Receding Hairline",
		"Scientifically Proven to Make You 17% Cooler",
		"Proudly Made By Monkeys on Typewriters",
		"100% Free of Artificial Intelligence (We Use Natural Stupidity)",
		"Server Hamsters Fed Regularly with Your Cookies",
		"Now Compatible with Your Toaster (Results May Vary)",
		"Caution: Website May Become Self-Aware and Start a Blog",
		"Warning: May Contain Traces of Actual Useful Information",
	}

	randomIndex := rand.Intn(len(phrases))

	return phrases[randomIndex]
}
