package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if data != nil {
		log.Println(data)
		json.NewEncoder(w).Encode(data)
	}
}
