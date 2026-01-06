package sharedjson

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, code, message string) {
	WriteJSON(w, status, contracts.APIResponse{
		Error: &contracts.APIError{
			Code:    code,
			Message: message,
		},
	})
}
