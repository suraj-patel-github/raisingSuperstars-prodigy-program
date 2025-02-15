package user

import (
	"context"
	"encoding/json"
	"net/http"

	"prodigy-program/service"
	httptransport "github.com/go-kit/kit/transport/http"

)

func NewHTTPHandler(svc service.UserService) http.Handler {
	mux := http.NewServeMux()

	registerUserHandler := httptransport.NewServer(
		MakeRegisterUserEndpoint(svc),
		decodeRegisterUserRequest,
		encodeResponse,
	)

	mux.Handle("/registerUser", registerUserHandler)
	return mux
}

// Decode HTTP request to struct
func decodeRegisterUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// Encode struct to HTTP response
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
