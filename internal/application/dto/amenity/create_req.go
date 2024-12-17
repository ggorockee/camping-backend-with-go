package amenitydto

type CreateAmenityReq struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
