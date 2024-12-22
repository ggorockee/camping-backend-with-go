package spotrepository

import (
	reviewdto "camping-backend-with-go/internal/application/dto/review"
	spotdto "camping-backend-with-go/internal/application/dto/spot"
	"camping-backend-with-go/internal/domain/entity"
	amenityrepository "camping-backend-with-go/internal/domain/repository/amenity"
	categoryrepository "camping-backend-with-go/internal/domain/repository/category"
	userrepository "camping-backend-with-go/internal/domain/repository/user"
	"fmt"

	"camping-backend-with-go/pkg/util"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type SpotRepository interface {
	CreateSpot(input *spotdto.CreateSpotReq, context ...*fiber.Ctx) (*entity.Spot, error)
	UpdateSpot(input *spotdto.UpdateSpotReq, id string, context ...*fiber.Ctx) (*entity.Spot, error)
	GetSpotById(id string, context ...*fiber.Ctx) (*entity.Spot, error)
	DeleteSpot(id string, context ...*fiber.Ctx) error
	GetAllSpots(context ...*fiber.Ctx) (*[]entity.Spot, error)
	GetReviewsFromSpot(spot *entity.Spot, context ...*fiber.Ctx) (*[]entity.Review, error)
	CreateSpotReview(input *reviewdto.CreateSpotReviewReq, spot *entity.Spot, context ...*fiber.Ctx) (*entity.Review, error)
}

type spotRepository struct {
	dbConn       *gorm.DB
	userRepo     userrepository.UserRepository
	categoryRepo categoryrepository.CategoryRepository
	amenityRepo  amenityrepository.AmenityRepository
}

func (r *spotRepository) CreateSpot(input *spotdto.CreateSpotReq, context ...*fiber.Ctx) (*entity.Spot, error) {
	// 컨텍스트 파싱
	c, err := util.ContextParser(context...)
	if err != nil {
		return nil, fmt.Errorf("context parsing failed: %w", err)
	}

	// 사용자 ID 토큰에서 추출
	userId := r.userRepo.GetValueFromToken("user_id", c)

	// 사용자 정보 조회
	owner, err := r.userRepo.GetUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Spot 엔티티 초기화
	var spot entity.Spot

	// input 데이터를 spot에 복사
	if err := copier.Copy(&spot, input); err != nil {
		return nil, err
	}

	spot.UserId = owner.Id // foreginKey

	// 카테고리 조회
	category, err := r.categoryRepo.GetCategoryById(*input.Category)
	if err != nil {
		return nil, err
	}

	spot.CategoryId = &category.Id

	// 트랜잭션 시작
	err = r.dbConn.Transaction(func(tx *gorm.DB) error {
		// Spot 생성
		if err := tx.Create(&spot).Error; err != nil {
			return fmt.Errorf("failed to create spot: %w", err)
		}

		// Amenities 처리
		if input.Amenities != nil {
			amenities := make([]entity.Amenity, 0)
			for _, amenityId := range *input.Amenities {
				amenity, err := r.amenityRepo.GetAmenityById(amenityId)
				if err != nil {
					return fmt.Errorf("failed to get amenity: %w", err)
				}
				amenities = append(amenities, *amenity)
			}

			// Amenities 연결
			if err := tx.Model(&spot).Association("Amenities").Replace(amenities); err != nil {
				return fmt.Errorf("failed to associate amenities: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &spot, nil
}

func (r *spotRepository) UpdateSpot(input *spotdto.UpdateSpotReq, id string, context ...*fiber.Ctx) (*entity.Spot, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)
	userId := r.userRepo.GetValueFromToken("user_id", c)
	spot, err := r.GetSpotById(id)
	if err != nil {
		return nil, err
	}

	if userId != spot.GetId() {
		return nil, errors.New("permission denied")
	}

	if err := copier.Copy(spot, input); err != nil {
		return nil, err
	}

	if err := r.dbConn.Save(spot).Error; err != nil {
		return nil, err
	}

	return spot, nil
}

func (r *spotRepository) GetSpotById(id string, context ...*fiber.Ctx) (*entity.Spot, error) {
	var spot entity.Spot

	err := r.dbConn.
		Preload("User").
		Preload("Amenities").
		Preload("Category").
		Preload("Reviews.User"). // Review의 User 정보도 함께 로드
		Where("id = ?", id).
		First(&spot).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("spot not found: %w", err)
		}
		return nil, fmt.Errorf("error fetching spot: %w", err)
	}

	return &spot, nil
}

func (r *spotRepository) DeleteSpot(id string, context ...*fiber.Ctx) error {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	userId := r.userRepo.GetValueFromToken("user_id", c)
	spot, err := r.GetSpotById(id)
	if err != nil {
		return err
	}

	if spot.GetId() != userId {
		return errors.New("permission denied")
	}

	if err := r.dbConn.Delete(spot).Error; err != nil {
		return err
	}

	return nil
}

func (r *spotRepository) GetAllSpots(context ...*fiber.Ctx) (*[]entity.Spot, error) {
	var spots []entity.Spot
	if err := r.dbConn.Preload("User").Find(&spots).Error; err != nil {
		return nil, err
	}

	return &spots, nil
}

func (r *spotRepository) GetReviewsFromSpot(spot *entity.Spot, context ...*fiber.Ctx) (*[]entity.Review, error) {
	var reviews []entity.Review

	if err := r.dbConn.Where("spot_id = ?", spot.GetId()).Preload("Spot").Find(&reviews).Error; err != nil {
		return nil, err
	}

	return &reviews, nil
}

func (r *spotRepository) CreateSpotReview(input *reviewdto.CreateSpotReviewReq, spot *entity.Spot, context ...*fiber.Ctx) (*entity.Review, error) {
	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	requestUser, ok := c.Locals("request_user").(entity.User)
	if !ok {
		return nil, errors.New("use is not authenticated")
	}

	// spot은 이미 불러와짐
	var review entity.Review

	review.UserId = requestUser.Id
	review.SpotId = &spot.Id

	if err := copier.Copy(&review, input); err != nil {
		return nil, err
	}

	// Spot이 이미 존재한다고 명시적으로 설정
	err = r.dbConn.Transaction(func(tx *gorm.DB) error {
		tx = tx.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
		if err := tx.Create(&review).Error; err != nil {
			return fmt.Errorf("error creating review: %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &review, nil
}

func NewSpotRepository(
	db *gorm.DB,
	u userrepository.UserRepository,
	c categoryrepository.CategoryRepository,
	a amenityrepository.AmenityRepository,
) SpotRepository {
	return &spotRepository{
		dbConn:       db,
		userRepo:     u,
		categoryRepo: c,
		amenityRepo:  a,
	}
}
