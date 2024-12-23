package adapter

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// FiberAdapter adapts an http.Handler to a Fiber handler.
func FiberAdapter(handler http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		r, err := fiberToStdRequest(c)
		if err != nil {
			return err
		}
		w := &fiberResponseWriter{ctx: c}
		handler.ServeHTTP(w, r)
		return nil
	}
}

// FiberAdapterFunc adapts an http.HandlerFunc to a Fiber handler.
func FiberAdapterFunc(handlerFunc http.HandlerFunc) fiber.Handler {
	return FiberAdapter(handlerFunc)
}
