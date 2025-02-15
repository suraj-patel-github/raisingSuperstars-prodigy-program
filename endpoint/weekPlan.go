package endpoint

import (
	"context"
	"prodigy-program/service"

	"github.com/go-kit/kit/endpoint"
)

type CreateWeekPlanRequest struct {
	Description string `json:"description"`
	UserID      int    `json:"userId"`
}

type CreateWeekPlanResponse struct {
	Code    int    `json:"code"`
	Success string `json:"success"`
	WeekID  int    `json:"week_id"`
}

type GetWeekPlanRequest struct {
	ID int `json:"id"`
}

func MakeCreateWeekPlanEndpoint(svc service.WeekPlanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateWeekPlanRequest)
		id, err := svc.CreateWeekPlan(req.Description, req.UserID)
		if err != nil {
			return nil, err
		}
		return CreateWeekPlanResponse{Code: 200, Success: "success", WeekID: id}, nil
	}
}
