package endpoint

import (
	"context"
	"prodigy-program/service"
	"prodigy-program/types"

	"github.com/go-kit/kit/endpoint"
)

type UpdateDayPlanRequest struct {
	DayNumber  int              `json:"dayNumber"`
	Activities []types.Activity `json:"activities"`
}

func MakeUpdateDayPlanEndpoint(svc service.DayPlanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateDayPlanRequest)
		err := svc.UpdateDayPlan(req.DayNumber, req.Activities)
		if err != nil {
			return nil, err
		}
		return map[string]string{"message": "Day plan updated successfully"}, nil
	}
}
