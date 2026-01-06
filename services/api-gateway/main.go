package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/env"
	sharedjson "ride-sharing/shared/json"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", HandleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var req contracts.PreviewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sharedjson.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
		return
	}

	body, _ := json.Marshal(req)

	resp, err := http.Post(
		"http://trip-service:8083/preview",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		sharedjson.WriteError(w, http.StatusBadGateway, "TRIP_SERVICE_UNAVAILABLE", err.Error())
		return
	}
	defer resp.Body.Close()

	// Se o trip-service retornou erro, propaga
	if resp.StatusCode >= 400 {
		sharedjson.WriteError(w, resp.StatusCode, "TRIP_SERVICE_ERROR", "failed to preview trip")
		return
	}

	// Decodifica resposta do trip-service
	var tripResponse any
	if err := json.NewDecoder(resp.Body).Decode(&tripResponse); err != nil {
		sharedjson.WriteError(w, http.StatusInternalServerError, "INVALID_RESPONSE", "invalid trip service response")
		return
	}

	// Envelopa no contrato do gateway
	sharedjson.WriteJSON(w, http.StatusOK, contracts.APIResponse{Data: tripResponse})
}
