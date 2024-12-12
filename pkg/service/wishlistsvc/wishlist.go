package wishlistsvc

import (
	"camping-backend-with-go/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetWishList(contexts ...*fiber.Ctx) (*[]entities.WishList, error)
}

type controller struct {
	repository Repository
}

// GetWishList implements Controller.
func (c *controller) GetWishList(contexts ...*fiber.Ctx) (*[]entities.WishList, error) {
	return c.repository.GetWishList(contexts...)
}

func NewController(r Repository) Controller {
	return &controller{repository: r}
}
