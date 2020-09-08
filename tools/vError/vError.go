package vError

import (
	"log"
	"net/http"
)

func WriteError(message string, statusCode int, err error, w http.ResponseWriter) {
	//w.WriteHeader(statusCode)
	//w.Write([]byte(message))
	http.Error(w, message, statusCode)
	log.Println(err)
}
