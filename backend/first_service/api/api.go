package api

import (
	"context"

	"first_service/dao"
	"first_service/proto"
)

type ServiceServer struct {
	proto.UnimplementedServiceServer
	Dao dao.DAO
}

func (s *ServiceServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	err := s.Dao.CreateMessage(req.Naam)
	if err != nil {
		return nil, err
	}
	return &proto.HelloResponse{Message: "Hello " + req.Naam}, nil
}
