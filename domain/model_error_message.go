package domain

import (
	"encoding/json"
)

//ErrorMessage is a wrapper type to return a JSON object with an error message
type ErrorMessage struct {
	Message string
}

//Bytes returns the ErrorMessage JSON bytes
func (errorMessage *ErrorMessage) Bytes() []byte {
	errorMessageJSON, err := json.Marshal(errorMessage)
	if err != nil {
		return nil
	}

	return errorMessageJSON
}

//GetErrorMessageBytes returns the bytes of the error message
func GetErrorMessageBytes(message string, err error) []byte {
	error := ErrorMessage{Message: message + err.Error()}
	return error.Bytes()
}
