package messages

// Success keywords
const INSERTED_SUCCESSFULLY = "Document has been created successfully"
const LOGGEDIN_SUCCESFULLY = "user logged-in succesfully"
const TOKEN_REFRESHED_SUCCESFULLY = "Token refreshed succesfully"

type successMessage struct {
	Success string
}

func GetSuccessMessage(message string) successMessage {
	return successMessage{message}
}
