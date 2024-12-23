package adapter

import (
	"bytes"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func fiberToStdRequest(c *fiber.Ctx) (*http.Request, error) {
	// creates a new http.Request
	url := c.Request().URI().String()
	body := bytes.NewReader(c.Request().Body())
	method := string(c.Request().Header.Method())
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// copy headers from Fiber to http.Request
	req.Header = make(http.Header)
	c.Request().Header.VisitAll(func(k, v []byte) {
		req.Header.Add(string(k), string(v))
	})

	// success
	return req, nil
}

// fiberResponseWriter is a custom http.ResponseWriter for Fiber.
type fiberResponseWriter struct {
	ctx      *fiber.Ctx
	headers  http.Header
	status   int
	response bytes.Buffer
}

func (w *fiberResponseWriter) Header() http.Header {
	if w.headers == nil {
		w.headers = make(http.Header)
	}
	return w.headers
}

func (w *fiberResponseWriter) Write(data []byte) (int, error) {
	return w.response.Write(data)
}

func (w *fiberResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ctx.Status(statusCode)
}
