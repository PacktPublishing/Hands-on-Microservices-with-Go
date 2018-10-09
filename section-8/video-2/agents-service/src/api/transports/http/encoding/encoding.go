package encoding

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/requests"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/responses"
)

var (
	// ManageID parameter is missing. Respond 404.
	ErrNoAgentID = errors.New("AgentID is required.")
	// AgentID was not a number. Respond 404
	ErrAgentIDNotNumber = errors.New("AgentID is not a number")
)

func DecodeInsertAgentPlayerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.InsertAgentPlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetAgentByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		return nil, ErrNoAgentID
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, ErrAgentIDNotNumber
	}

	return requests.GetAgentByIDRequest{AgentID: uint32(id)}, nil
}

func EncodeGetAgentByIDResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	res, ok := response.(responses.GetAgentByIDResponse)
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

func EncodeInsertAgentPlayerResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	res, ok := response.(responses.InsertAgentPlayerResponse)
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
