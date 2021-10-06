package main

import (
	"context"
	"fmt"
	"github.com/anantadwi13/belajar-go/grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:4321", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := model.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetAll(ctx, &model.GetUserByIdReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetUsers())
}

func main2() {
	user := &model.User{
		Id:          123,
		Name:        "YUDA",
		Phones:      []*model.User_PhoneNumber{{Number: "123", Type: model.User_HOME}},
		UserType:    model.User_ADMIN,
		LastUpdated: timestamppb.Now(),
	}

	data, err := proto.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Println(user)
}
