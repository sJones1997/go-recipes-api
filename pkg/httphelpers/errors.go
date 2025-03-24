package httphelpers

import (
	"net/http"
)

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request){
	http.Error(w, "Not Found", http.StatusNotFound)
}