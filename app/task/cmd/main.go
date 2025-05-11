package main

import (
	"github.com/whuyfk/godolist/app/task/internal/service"
	pb "github.com/whuyfk/godolist/idl/pb/task"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterTaskServiceServer(server, service.GetTaskSrv())
	err = server.Serve(lis)
	if err != nil {
		panic(err)
		return
	}
}
