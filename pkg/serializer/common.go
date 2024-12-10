package serializer

import (
	"github.com/gofiber/fiber/v2"
)

func MakeContext(contexts []*fiber.Ctx) *fiber.Ctx {
	if len(contexts) > 0 {
		return contexts[0]
	}
	return nil
}
