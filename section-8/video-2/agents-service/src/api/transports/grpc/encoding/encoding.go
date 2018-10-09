package encoding

import (
	"context"
	"errors"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/grpc/pb"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/requests"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/responses"
)

func DecodeInsertAgentPlayerRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.InsertAgentPlayerRequest)
	return requests.InsertAgentPlayerRequest{AgentID: uint32(req.AgentID), PlayerID: uint32(req.PlayerID)}, nil
}

func DecodeGetAgentByIDRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetAgentByIDRequest)
	return requests.GetAgentByIDRequest{AgentID: uint32(req.AgentID)}, nil
}

func EncodeInsertAgentPlayerResponse(_ context.Context, response interface{}) (interface{}, error) {
	pbReply := &pb.InsertAgentPlayerReply{}

	res, ok := response.(responses.InsertAgentPlayerResponse)
	if !ok {
		return nil, errors.New("Error when casting response.")
	}
	if res.Err != "" {
		return nil, errors.New(res.Err)
	}
	pbReply.Success = true

	return pbReply, nil
}

func EncodeGetAgentByIDResponse(_ context.Context, response interface{}) (interface{}, error) {

	pbReply := &pb.GetAgentByIDReply{}

	res, ok := response.(responses.GetAgentByIDResponse)
	if !ok {
		return nil, errors.New("Error when casting response.")
	}
	if res.Err != "" {
		return nil, errors.New(res.Err)
	}
	pbReply.AgentID = res.Agent.ID
	pbReply.Account = res.Agent.Account
	pbReply.Name = res.Agent.Name
	return pbReply, nil
}
