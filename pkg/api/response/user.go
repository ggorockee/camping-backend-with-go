package response

// entity.user -> tinyUserResponse로 -> marshalJSON()으로

type TinyUserReponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
