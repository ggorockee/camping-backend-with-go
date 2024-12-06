package presenter

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewJsonResponse(error bool, message string, data any) JsonResponse {
	if data == nil {
		data = []interface{}{} // nil일 경우 빈 슬라이스로 초기화
	}

	return JsonResponse{
		Error:   error,
		Message: message,
		Data:    data,
	}
}
