package messages

// Error keywords
const CONNECTION_ERROR = "DataBase is not connected"
const BAD_REQUEST = "Form data is invalid"

type errorMessage struct {
	Error string
}

func GetErrorMessage(message string) errorMessage {
	return errorMessage{message}
}
