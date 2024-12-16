package helpers

import (
	entities2 "camping-backend-with-go/internal/domain"
	"camping-backend-with-go/pkg/dto"
	"camping-backend-with-go/pkg/serializer"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SerializeSpot(spots []entities2.Spot, db *gorm.DB, c *fiber.Ctx) []dto.SpotListOut {
	spotListRes := make([]dto.SpotListOut, 0)
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

func SerializeUser(user entities2.User) dto.TinyUserOut {
	userSerializer := serializer.NewUserSerializer(&user)
	return userSerializer.TinyUserSerialize()
}
