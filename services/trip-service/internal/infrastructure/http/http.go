package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/contracts"
	sharedjson "ride-sharing/shared/json"
)

type HandlerService struct {
	Service domain.TripService
}

func (h *HandlerService) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody contracts.PreviewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parsed JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	t, err := h.Service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Println(err)
	}

	sharedjson.WriteJSON(w, http.StatusOK, t)
}
