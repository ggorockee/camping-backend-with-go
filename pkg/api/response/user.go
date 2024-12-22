package response

import (
	"camping-backend-with-go/internal/domain/entity"
	"encoding/json"
)

// entity.user -> tinyUserResponse로 -> marshalJSON()으로

type TinyUserReponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type TinyUserAdapter struct {
	*TinyUserReponse
}

func NewTinyUserAdapter(u *entity.User) *TinyUserAdapter {
	user := &TinyUserReponse{
		Id:    u.Id,
		Email: u.Email,
	}
	return &TinyUserAdapter{
		TinyUserReponse: user,
	}

}

func (a *TinyUserAdapter) MarshalJSON() ([]byte, error) {
	userResponse := TinyUserReponse{
		Id:    a.Id,
		Email: a.Email,
	}
	return json.Marshal(userResponse)
}
