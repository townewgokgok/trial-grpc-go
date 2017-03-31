package main

import (
	"net"

	"github.com/townewgokgok/trial-grpc-go/model"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type trialGrpcServerImpl struct {
	people []*model.Person
}

func (s *trialGrpcServerImpl) addPerson(person *model.Person) {
	person.Id = int32(len(s.people) + 1)
	s.people = append(s.people, person)
}

func (s *trialGrpcServerImpl) GetPerson(ctx context.Context, req *model.IdRequest) (*model.Person, error) {
	return s.people[req.Id-1], nil
}

func newTrialGrpcServerImpl() *trialGrpcServerImpl {
	s := new(trialGrpcServerImpl)
	s.addPerson(&model.Person{
		Name:  "townewgokgok",
		Email: "townewgokgok@gmail.com",
		Phone: []*model.Person_PhoneNumber{
			{Number: "03-3333-3333", Type: model.Person_HOME},
			{Number: "080-8888-8888", Type: model.Person_MOBILE},
		},
	})
	s.addPerson(&model.Person{
		Name:  "hogehoge",
		Email: "hogehoge@example.com",
		Phone: []*model.Person_PhoneNumber{},
	})
	return s
}

func main() {
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	s := newTrialGrpcServerImpl()
	model.RegisterTrialGrpcServer(grpcServer, s)
	grpcServer.Serve(listener)
}
