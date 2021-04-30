package helper

type Message struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ErrorMessage(message string) Message {
	return Message{
		Status:  false,
		Message: message,
	}
}

func SuccessMessage(message string) Message {
	return Message{
		Status:  true,
		Message: message,
	}
}
