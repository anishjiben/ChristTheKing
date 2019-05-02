package messages

// Error keywords
const CONNECTION_ERROR = "DataBase is not connected"
const INSERTION_FAILED = "Document insertion failed"
const BAD_REQUEST = "Form data is invalid"
const USER_EXIST = "User exist"
const UN_AUTHORIZED_USER = "User is unauthorized"

type errorMessage struct {
	Error string
}

func GetErrorMessage(message string) errorMessage {
	return errorMessage{message}
}
