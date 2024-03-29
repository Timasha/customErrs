package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (e *ErrsParser) CatchCustomErrMiddleware(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {
		customErr, isInternal := e.Parse(err.Error())
		if isInternal && customErr.HttpCode < 500 {
			return customErr
		}
	}

	return err
}
