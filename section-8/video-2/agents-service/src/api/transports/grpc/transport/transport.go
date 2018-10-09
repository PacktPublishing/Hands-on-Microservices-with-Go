package transport

import (
	"context"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/grpc/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type GrpcServer struct {
	InsertAgentPlayerHandler grpctransport.Handler
	GetAgentByIDHandler      grpctransport.Handler
}

func (s *GrpcServer) InsertAgentPlayer(ctx context.Context, req *pb.InsertAgentPlayerRequest) (*pb.InsertAgentPlayerReply, error) {
	_, rep, err := s.InsertAgentPlayerHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.InsertAgentPlayerReply), nil
}
func (s *GrpcServer) GetAgentByID(ctx context.Context, req *pb.GetAgentByIDRequest) (*pb.GetAgentByIDReply, error) {
	_, rep, err := s.GetAgentByIDHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAgentByIDReply), nil
}
