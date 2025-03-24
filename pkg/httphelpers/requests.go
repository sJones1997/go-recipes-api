package httphelpers

import (
	"encoding/json"
	"net/http"
)

func DecodeBody[T any](w http.ResponseWriter, r *http.Request) (T, error){
	var value T
	if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
		InternalServerErrorHandler(w, r)
		return value, err
	}
	return value, nil
}