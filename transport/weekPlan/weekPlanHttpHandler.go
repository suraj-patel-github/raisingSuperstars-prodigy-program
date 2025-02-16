package weekPlan

import (
	"context"
	"encoding/json"
	"net/http"
	"prodigy-program/endpoint"
	"prodigy-program/service"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewWeekPlanHandler(svc service.WeekPlanService) http.Handler {
	mux := http.NewServeMux()

	createWeekPlanHandler := httptransport.NewServer(
		endpoint.MakeCreateWeekPlanEndpoint(svc),
		decodeCreateWeekPlanRequest,
		encodeResponse,
	)

	getWeekPlanHandler := httptransport.NewServer(
		endpoint.MakeGetWeekPlanByWeekIdEndpoint(svc),
		decodeGetWeekPlanRequest,
		encodeResponse,
	)

	mux.Handle("/weekplan", createWeekPlanHandler)
	mux.Handle("/getweekplan", getWeekPlanHandler)
	return mux
}

func decodeCreateWeekPlanRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateWeekPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// Decode request for getting a week plan by ID
func decodeGetWeekPlanRequest(_ context.Context, r *http.Request) (interface{}, error) {

	req := endpoint.GetWeekPlanRequest{}
	queryParams := r.URL.Query()
	weekIDStr := queryParams.Get("weekId")
	if weekIDStr == "" {
		return nil, http.ErrMissingFile // Return an error if week_id is missing
	}
	weekID, err := strconv.Atoi(weekIDStr)
	if err != nil {
		return nil, err
	}
	req.ID = weekID

	dayNumberStr := queryParams.Get("dayNumber")
	if dayNumberStr != "" {
		dayNUmber, err := strconv.Atoi(dayNumberStr)
		if err != nil {
			return nil, err
		}
		req.DayNumber = dayNUmber
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
