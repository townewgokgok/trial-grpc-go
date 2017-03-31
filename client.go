package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/townewgokgok/trial-grpc-go/model"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5555", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := model.NewTrialGrpcClient(conn)
	for i := 1; i <= 2; i++ {
		req := model.IdRequest{Id: int32(i)}
		res, err := client.GetPerson(context.Background(), &req)
		if err != nil {
			grpclog.Fatalf("fail to GetPerson: %v", err)
		}
		fmt.Printf("%d: %s <%s>\n", res.Id, res.Name, res.Email)
		for _, phone := range res.Phone {
			fmt.Printf("    %s: %s\n", phone.Type, phone.Number)
		}
		if len(res.Phone) == 0 {
			fmt.Println("    (no phone)")
		}
	}
}
