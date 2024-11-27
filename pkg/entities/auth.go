package entities

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
