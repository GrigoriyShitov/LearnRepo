// nolint:golines
package server

import (
	"net/http"

	"git.itmo.su/go-modules/xerr/v4"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
	ErrorCode    *string `json:"error_code"`
	ErrorMessage *string `json:"error_message" extensions:"x-nullable"`
	Data         any     `json:"data" extensions:"x-nullable"`
}

func (s *Server) response(c *fiber.Ctx, data any, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	if err == nil {
		return c.Status(http.StatusOK).JSON(Body{Data: data})
	}
	if xe, ok := xerr.As(err); ok {
		msg := xe.Localize(c.UserContext())
		return c.Status(xe.Kind.HTTPCode).JSON(Body{
			ErrorCode:    &xe.Kind.Key,
			ErrorMessage: &msg,
			Data:         data,
		})
	}

	unknownErr := xerr.ErrInternal
	msg := unknownErr.Localize(c.UserContext())
	return c.Status(unknownErr.Kind.HTTPCode).JSON(Body{
		ErrorCode:    &unknownErr.Kind.Key,
		ErrorMessage: &msg,
	})
}
