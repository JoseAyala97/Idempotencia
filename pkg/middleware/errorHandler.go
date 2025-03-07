package middleware

import (
	"IdEmpotencia/pkg/apperror"
	"IdEmpotencia/pkg/message"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ErrorHandler(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {
		if e, ok := err.(*apperror.AppError); ok {
			log.Error(message.ErrorResponse(e))
			return c.Status(e.Code).JSON(message.ErrorResponse(e))
		}
		log.Error(message.ErrorResponse(err))
		return c.Status(fiber.StatusInternalServerError).JSON(message.ErrorResponse(err))
	}
	return nil
}
