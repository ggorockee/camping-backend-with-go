package amenitydto

type UpdateAmenityReq struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
