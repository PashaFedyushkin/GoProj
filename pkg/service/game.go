package service

import (
	"context"
	"log"
	"math/rand"

	pb "github.com/PashaFedyushkin/GoProj/pkg/api"
	"google.golang.org/grpc"
)

// GAMEServer ...
type GAMEServer struct{ pb.UnimplementedGameServer }

// Game ...
func (s *GAMEServer) Game(ctx context.Context, req *pb.GameRequest) (*pb.GameResponse, error) {
	var a int
	a = rand.Intn(6)
	if req.GetReq() == int32(a) {
		return &pb.GameResponse{Res: "You win"}, nil
	}
	return &pb.GameResponse{Res: "You lose"}, nil

	//type GameService struct{}

	//func NewGameService() *AuthService {}
}

type GameService struct {
	port string
}

func NewGameService(port string) *GameService {
	return &GameService{port: port}
}

func (g *GameService) LetsPlay(number int32) (string, error) {
	conn, err := grpc.Dial(g.port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := pb.NewGameClient(conn)
	res, err := c.Game(context.Background(), &pb.GameRequest{Req: number})
	//if err != nil {
	//	log.Fatal(err)
	//}
	return res.GetRes(), err
}
