package responses

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
}
