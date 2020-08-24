package transports

import (
	"context"
	"encoding/json"
	"net/http"
)

type getRequest struct {
}

type getResponse struct {
	Date string `json:"date"`
	Err  string `json:"err, omitempty"`
}

type StatusRequest struct {
}

type StatusResponse struct {
	Status string `json:"status"`
}

func DecodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req StatusRequest
	return req, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
