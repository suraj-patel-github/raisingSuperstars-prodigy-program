package dayplan

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"prodigy-program/endpoint"
	"prodigy-program/service"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewDayPlanHandler(svc service.DayPlanService) http.Handler {
	mux := http.NewServeMux()

	updateDayPlanHandler := httptransport.NewServer(
		endpoint.MakeUpdateDayPlanEndpoint(svc),
		decodeDayPlanRequest,
		encodeResponse,
	)

	mux.Handle("/updatedayplan", updateDayPlanHandler)
	return mux
}

func decodeDayPlanRequest(_ context.Context, r *http.Request) (interface{}, error) {

	req := endpoint.UpdateDayPlanRequest{}
	// Get dayID from query params
	dayNumberStr := r.URL.Query().Get("dayNumber")
	dayNumber, err := strconv.Atoi(dayNumberStr)
	if err != nil {
		return nil, fmt.Errorf("invalid day number")
	}
	req.DayNumber = dayNumber

	// will have array of activies in the request body
	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req.Activities); err != nil {
		return nil, err
	}

	return req, nil
}


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
