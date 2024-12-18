package response

import (
	"time"
)

type ReviewTinyRes struct {
	Id   uint        `json:"id"`
	User UserTinyRes `json:"user"`

	Spot SpotTinyRes `json:"spot"`

	Payload string `json:"payload"`
	Rating  int    `json:"rating"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
