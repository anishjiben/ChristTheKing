package messages

// Success keywords
const INSERTED_SUCCESSFULLY = "Document has been created successfully"
const LOGGEDIN_SUCCESFULLY = "user logged-in succesfully"

type successMessage struct {
	Success string
}

func GetSuccessMessage(message string) successMessage {
	return successMessage{message}
}
