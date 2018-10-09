package endpoints

import (
	"context"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/service"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/requests"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/responses"

	"github.com/go-kit/kit/endpoint"
)

func MakeInsertAgentPlayerEndpoint(svc service.AgentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.InsertAgentPlayerRequest)
		err := svc.InsertAgentPlayer(req.AgentID, req.PlayerID)
		if err != nil {
			return responses.InsertAgentPlayerResponse{Err: err.Error()}, nil
		}
		return responses.InsertAgentPlayerResponse{Err: ""}, nil
	}
}

func MakeGetAgentByIDEndpoint(svc service.AgentsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.GetAgentByIDRequest)
		v, err := svc.GetAgentByID(req.AgentID)
		if err != nil {
			return responses.GetAgentByIDResponse{Agent: nil, Err: err.Error()}, nil
		}
		return responses.GetAgentByIDResponse{Agent: v, Err: ""}, nil
	}
}
