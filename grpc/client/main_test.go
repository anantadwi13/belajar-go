package main

import (
	"context"
	"github.com/anantadwi13/belajar-go/grpc/model"
	"google.golang.org/grpc"
	"testing"
)

func BenchmarkClient(b *testing.B) {
	conn, err := grpc.Dial("localhost:4321", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := model.NewUserServiceClient(conn)

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		res, err := c.GetAll(ctx, &model.GetUserByIdReq{})
		if err != nil {
			b.Error(err)
		}
		if len(res.GetUsers()) != 1 || res.GetUsers()[0].Name != "YUDA" {
			b.Error("hasil tidak cocok")
		}
	}
}
