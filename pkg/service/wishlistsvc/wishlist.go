package wishlistsvc

import (
	"camping-backend-with-go/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetWishLists(contexts ...*fiber.Ctx) (*[]entities.WishList, error)
	WishListToggle(contexts ...*fiber.Ctx) (*entities.WishList, error)
}

type controller struct {
	repository Repository
}

func (c *controller) WishListToggle(contexts ...*fiber.Ctx) (*entities.WishList, error) {
	//TODO implement me
	panic("implement me")
}

// GetWishList implements Controller.
func (c *controller) GetWishLists(contexts ...*fiber.Ctx) (*[]entities.WishList, error) {
	return c.repository.GetWishList(contexts...)
}

func NewController(r Repository) Controller {
	return &controller{repository: r}
}
