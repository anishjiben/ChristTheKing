package messages

// Success keywords
const INSERTED_SUCCESSFULLY = "Document has been created successfully"
const LOGGEDIN_SUCCESFULLY = "user logged-in succesfully"
const LOGGEDOUT_SUCCESFULLY = "user logged-out succesfully"
const TOKEN_REFRESHED_SUCCESFULLY = "Token refreshed succesfully"
const UPDATED_SUCCESFULLY = "Document has been updated successfully"
const DELETED_SUCCESFULLY = "Document has been deleted succesfully"

type successMessage struct {
	Success string
}

func GetSuccessMessage(message string) successMessage {
	return successMessage{message}
}
