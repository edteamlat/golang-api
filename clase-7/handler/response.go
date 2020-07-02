package handler

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(messageType string, message string, data interface{}) response {
	return response{
		messageType,
		message,
		data,
	}
}
