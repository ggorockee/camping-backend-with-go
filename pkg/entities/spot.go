package entities

import (
	"time"
)

type Spot struct {
	Id         uint `json:"id" gorm:"primaryKey"`
	UserId     uint `json:"user_id"`
	User       User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	CategoryId *int `gorm:"default:null" json:"category_id"` // CategoryId가 null일 수가 있음
	// sqlite에서 SET NULL, mysql, postgresql에서는 SetNull
	// 배포시 아래 주석
	// sqlite 설정
	//Category Category `gorm:"foreignKey:CategoryId;constraint:OnDelete:SET NULL;"`

	// 배포시 아래 주석해제
	// rds 설정
	Category  Category  `gorm:"foreignKey:CategoryId;constraint:OnDelete:SetNull;"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	Author    string    `json:"author"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CoverImg string `json:"cover_img"`
}

type SpotSerializer interface {
	ListSerialize() SpotListOutputSchema
	DetailSerialize() SpotDetailOutputSchema
}

type spotSerializer struct {
	spot *Spot
	user UserSerializer
}

func (s *spotSerializer) ListSerialize() SpotListOutputSchema {

	return SpotListOutputSchema{
		Id:        int(s.spot.Id),
		User:      s.user.TinyUserSerialize(),
		Title:     s.spot.Title,
		Location:  s.spot.Location,
		Author:    s.spot.Author,
		CreatedAt: s.spot.CreatedAt,
		UpdatedAt: s.spot.UpdatedAt,
		Review:    s.spot.Review,
	}
}

func (s *spotSerializer) DetailSerialize() SpotDetailOutputSchema {
	return SpotDetailOutputSchema{
		Id:        int(s.spot.Id),
		User:      s.user.TinyUserSerialize(),
		Title:     s.spot.Title,
		Location:  s.spot.Location,
		Author:    s.spot.Author,
		CreatedAt: s.spot.CreatedAt,
		UpdatedAt: s.spot.UpdatedAt,
		Review:    s.spot.Review,
	}
}

func NewSpotSerializer(s *Spot, u UserSerializer) SpotSerializer {
	return &spotSerializer{spot: s, user: u}
}

// ============= input schema =============

type CreateSpotInputSchema struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

type UpdateSpotSchema struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	Review   string `json:"review"`
}

// ============= output schema =============

type SpotListOutputSchema struct {
	Id        int                  `json:"id"`
	User      TinyUserOutputSchema `json:"user"`
	Title     string               `json:"title"`
	Location  string               `json:"location"`
	Author    string               `json:"author"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Review    string               `json:"review"`
}

type SpotDetailOutputSchema struct {
	Id        int                  `json:"id"`
	User      TinyUserOutputSchema `json:"user"`
	Title     string               `json:"title"`
	Location  string               `json:"location"`
	Author    string               `json:"author"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Review    string               `json:"review"`
}
