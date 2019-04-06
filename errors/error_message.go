package errors

// Data base connection error keywords
const CONNECTION_ERROR = "DataBase is not connected"

type errorMessage struct {
	Error string
}

func GetErrorMessage(message string) errorMessage {
	return errorMessage{message}
}
