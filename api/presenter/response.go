package presenter

type JsonResponse struct {
	Status bool   `json:"status"`
	Data   any    `json:"data"`
	Error  string `json:"error"`
}
