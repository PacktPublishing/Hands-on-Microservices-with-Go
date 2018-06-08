package WTAServer

import (
	"context"
	"io"
	"log"

	pb "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/proto"
	repo "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/repository"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type MyWTAServer struct {
	Repo *repo.PsqlRepo
}

func (s *MyWTAServer) GetPlayerWithHighestRanking(ctx context.Context, in *pb.PlayerIdRequest) (*pb.PlayerWithRanking, error) {
	result, err := s.Repo.GetPlayerPlusHighestRanking(in.GetPlayerId())
	if err != nil {
		return nil, err
	}
	out := &pb.PlayerWithRanking{}
	out.Player = &pb.Player{}
	out.Ranking = &pb.Ranking{}

	out.Player.Id = result.ID
	out.Player.FirstName = result.FirstName
	out.Player.LastName = result.LastName

	out.Player.BirthDate = &timestamp.Timestamp{Seconds: int64(result.BirthDate.Unix()), Nanos: int32(0)}

	out.Player.IsRightHanded = result.IsRightHanded
	out.Player.CountryCode = result.CountryCode
	out.Ranking.Ranking = uint32(result.Ranking.RankingNumber)
	out.Ranking.RankingDate = &timestamp.Timestamp{Seconds: int64(result.RankingDate.Unix()), Nanos: int32(0)}
	out.Ranking.RankingPoints = float32(result.RankingPoints)
	out.Ranking.PlayerId = result.ID
	return out, nil
}

func (s *MyWTAServer) GetRankingsByPlayerId(in *pb.PlayerIdRequest, stream pb.WTA_GetRankingsByPlayerIdServer) error {
	rankings, err := s.Repo.GetRankingsByPlayerID(in.GetPlayerId())
	if err != nil {
		return err
	}
	for _, r := range rankings {
		res := &pb.Ranking{}
		res.PlayerId = r.PlayerID
		res.Ranking = uint32(r.RankingNumber)
		res.RankingDate = &timestamp.Timestamp{Seconds: int64(r.RankingDate.Unix()), Nanos: int32(0)}
		res.RankingPoints = float32(r.RankingPoints)
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	//Sending nil we have finished writing responses
	return nil
}

//stream is read and write
func (s *MyWTAServer) GetPlayers(stream pb.WTA_GetPlayersServer) error {
	results := &pb.PlayersReply{}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(results)
		}
		if err != nil {
			return err
		}

		log.Printf("Recieved request for PlayerID: &d", in.GetPlayerId())

		res, err := s.Repo.GetPlayer(in.GetPlayerId())
		if err != nil {
			return err
		}
		p := &pb.Player{}
		p.Id = res.ID
		p.FirstName = res.FirstName
		p.LastName = res.LastName
		p.IsRightHanded = res.IsRightHanded
		p.CountryCode = res.CountryCode
		p.BirthDate = &timestamp.Timestamp{Seconds: int64(res.BirthDate.Unix()), Nanos: int32(0)}
		results.Player = append(results.Player, p)
	}
	return nil
}
