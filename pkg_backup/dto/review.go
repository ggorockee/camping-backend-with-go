package dto

type CreateSpotReviewReq struct {
	Payload string `json:"payload"`
	Rating  int    `json:"rating"`
}
