package server

import "encoding/json"

type ErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type SuccessMessage struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

var defaultError = []byte("{\"code\":-1,\"message\":\"internal server error\"}")

func createError(code int, message string) []byte {
	var data = ErrorMessage{
		Code:    code,
		Message: message,
	}
	var jsonData []byte
	var err error

	jsonData, err = json.Marshal(data)
	if err != nil {
		return defaultError
	}

	return jsonData
}
