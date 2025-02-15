package weekPlan

import (
	"context"
	"encoding/json"
	"net/http"
	"prodigy-program/endpoint"
	"prodigy-program/service"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewWeekPlanHandler(svc service.WeekPlanService) http.Handler {
	mux := http.NewServeMux()

	createWeekPlanHandler := httptransport.NewServer(
		endpoint.MakeCreateWeekPlanEndpoint(svc),
		decodeCreateWeekPlanRequest,
		encodeResponse,
	)

	mux.Handle("/weekplan", createWeekPlanHandler)
	return mux
}

func decodeCreateWeekPlanRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateWeekPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
