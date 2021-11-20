package responses

type MessageResponse struct {
	Message string `json:"message"`
}

func OK() *MessageResponse {
	return &MessageResponse{Message: "OK"}
}

type ErrorResponse struct {
	Error string `json:"error"`
}
