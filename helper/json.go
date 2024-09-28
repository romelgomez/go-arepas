package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	ErrorPanic(err)
}

func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	ErrorPanic(err)
}

func WriteErrorResponse(writer http.ResponseWriter, message string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(map[string]string{
		"error": message,
	})
}
