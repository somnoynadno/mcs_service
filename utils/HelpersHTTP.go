package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}

func RespondJSON(w http.ResponseWriter, data []byte) {
	_, _ = w.Write(data)
}

var HandleOptions = func(w http.ResponseWriter, r *http.Request) {
	// nothing to say
	w.WriteHeader(http.StatusOK)
}