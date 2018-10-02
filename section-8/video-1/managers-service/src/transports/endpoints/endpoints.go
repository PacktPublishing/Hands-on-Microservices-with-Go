package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/managers-service/src/service"
	"github.com/gorilla/mux"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/managers-service/src/transports/requests"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/managers-service/src/transports/responses"

	"github.com/go-kit/kit/endpoint"
)

var (
	// ManageID parameter is missing. Respond 404.
	ErrNoManagerID = errors.New("ManagerID is required.")
	// ManagerID was not a number. Respond 404
	ErrManagerIDNotNumber = errors.New("ManagerID is not a number")
)

func MakeInsertManagerPlayerEndpoint(svc service.ManagersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.InsertManagerPlayerRequest)
		err := svc.InsertManagerPlayer(req.ManagerID, req.PlayerID)
		if err != nil {
			return responses.InsertManagerPlayerResponse{Err: err.Error()}, nil
		}
		return responses.InsertManagerPlayerResponse{Err: ""}, nil
	}
}

func MakeGetManagerByIDEndpoint(svc service.ManagersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.GetManagerByIDRequest)
		v, err := svc.GetManagerByID(req.ManagerID)
		if err != nil {
			return responses.GetManagerByIDResponse{Manager: nil, Err: err.Error()}, nil
		}
		return responses.GetManagerByIDResponse{Manager: v, Err: ""}, nil
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
	idstr, ok := vars["id"]
	if !ok {
		return nil, ErrNoManagerID
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, ErrManagerIDNotNumber
	}

	return requests.GetManagerByIDRequest{ManagerID: uint32(id)}, nil
}

func EncodeGetManagerByIDResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	res, ok := response.(responses.GetManagerByIDResponse)
	if !ok {
		w.WriteHeader(500)
		return errors.New("Error when casting response.")
	}
	if res.Err != "" {
		w.WriteHeader(500)
		return errors.New(res.Err)
	}
	str, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		return err
	}
	w.Write(str)
	return nil
}

func EncodeInsertManagerPlayerResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	res, ok := response.(responses.InsertManagerPlayerResponse)
	if !ok {
		w.WriteHeader(500)
		return errors.New("Error when casting response.")
	}
	if res.Err != "" {
		w.WriteHeader(500)
		return errors.New(res.Err)
	}
	w.WriteHeader(201)
	w.Write([]byte("Write was succesful."))
	return nil
}
