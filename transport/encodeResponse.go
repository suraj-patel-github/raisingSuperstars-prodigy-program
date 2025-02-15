package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

// Standard Response Encoder for Go-Kit
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}