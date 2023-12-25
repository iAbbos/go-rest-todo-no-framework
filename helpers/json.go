package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decooder := json.NewDecoder(r.Body)
	err := decooder.Decode(result)
	PanicIfError(err)
}

func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	PanicIfError(err)
}
