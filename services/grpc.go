package services

// here's for grpc handler
// same with rest.go, in here only for handler and call logic to specific service you want
// so the service is apart from handler
// you can run both of grpc and rest

import (
	"context"
	"encoding/json"

	pb "github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/proto"
)

type GrpcServer struct {
	pb.UnimplementedGrpcServerServer
	srvHello *Hello
}

func (grpc *GrpcServer) RegisterHello(acd *Hello) {
	grpc.srvHello = acd
}

func (grpc *GrpcServer) HelloWorld(ctx context.Context, request *pb.RequestHello) (*pb.ResponseHello, error) {
	getResp, _ := grpc.srvHello.SaidHello(&requestHello{
		Name: request.Name,
	})

	getAny, _ := json.Marshal(getResp.Data)
	return &pb.ResponseHello{Code: getResp.Code, Message: getResp.Message, Data: getAny}, nil
}
