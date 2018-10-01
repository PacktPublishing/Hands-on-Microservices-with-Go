package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/service"
	"github.com/gorilla/mux"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/transports/requests"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/transports/responses"

	"github.com/go-kit/kit/endpoint"
)

var (
	// ManageID parameter is missing. Respond 404.
	ErrNoManagerID = errors.New("ManagerID is required.")
)

func MakeInsertManagerPlayerEndpoint(svc service.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.InsertManagerPlayerRequest)
		v, err := svc.InsertManagerPlayerRequest(req.ManagerID, req.PlayerID)
		if err != nil {
			return responses.InsertManagerPlayerResponse{err.Error()}, nil
		}
		return responses.InsertManagerPlayerResponse{""}, nil
	}
}

func MakeGetManagerByIDEndpoint(svc service.ManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.GetManagerByIDRequest)
		v := svc.GetManagerByID(req.ManagerID)
		return responses.GetManagerByIDResponse{v}, nil
	}
}

func DecodeInsertManagerPlayerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.InsertManagerPlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetManagerByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrNoManagerID
	}
	return responses.GetManagerIDResponse{ManagerID: id}, nil
}

func DecodeInsertManagerPlayerResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response responses.InsertManagerPlayerResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeGetManagerByIDResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response responses.GetManagerByIDResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
