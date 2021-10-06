package main

import (
	"context"
	"github.com/anantadwi13/belajar-go/grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

type server struct {
	model.UnimplementedUserServiceServer
}

func (s *server) GetAll(ctx context.Context, req *model.GetUserByIdReq) (*model.GetUserByIdRes, error) {
	return &model.GetUserByIdRes{
		Users: []*model.User{
			{
				Id:          123,
				Name:        "YUDA",
				Phones:      []*model.User_PhoneNumber{{Number: "123", Type: model.User_HOME}},
				UserType:    model.User_ADMIN,
				LastUpdated: timestamppb.Now(),
			},
		},
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":4321")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	model.RegisterUserServiceServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
