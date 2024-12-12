package middleware

import "github.com/gofiber/fiber/v2"

func ContextParser(contexts ...*fiber.Ctx) *fiber.Ctx {
	if len(contexts) == 0 {
		return nil
	}

	return contexts[0]
}
