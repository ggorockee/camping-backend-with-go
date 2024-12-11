package util

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ContextParser(contexts ...*fiber.Ctx) (*fiber.Ctx, error) {
	if len(contexts) != 1 {
		return nil, errors.New("exactly one context must be provided")
	}
	return contexts[0], nil
}
