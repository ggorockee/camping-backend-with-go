package helpers

import (
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/serializer"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SerializeSpot(spots []entities.Spot, db *gorm.DB, c *fiber.Ctx) []dto.SpotListOut {
	var spotListRes []dto.SpotListOut
	for _, spot := range spots {
		userSerializer := serializer.NewUserSerializer(&spot.User)
		categorySerializer := serializer.NewCategorySerializer(&spot.Category)
		spotSerializer := serializer.NewSpotSerializer(&spot, userSerializer, categorySerializer)
		spotListOut := spotSerializer.ListSerialize(db, c)

		// User 필드가 비어있는지 확인
		if spotListOut.User != nil && spotListOut.User.Id == 0 {
			spotListOut.User = nil
		}

		spotListRes = append(spotListRes, spotListOut)
	}
	return spotListRes
}

func SerializeUser(user entities.User) dto.TinyUserOut {
	userSerializer := serializer.NewUserSerializer(&user)
	return userSerializer.TinyUserSerialize()
}
