package user

import (
	"context"
	"prodigy-program/service"

	"github.com/go-kit/kit/endpoint"
)

type RegisterUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response Structure
type RegisterUserResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
}

// Endpoint to Register a User
func MakeRegisterUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterUserRequest)
		userID, err := svc.RegisterUser(ctx, req.Name, req.Email)
		if err != nil {
			return nil, err
		}
		return RegisterUserResponse{Code: 201, Status: "success", UserID: userID}, nil
	}
}

