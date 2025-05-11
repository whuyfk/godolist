package rpc

import (
	"context"
	pb "github.com/whuyfk/godolist/idl/pb/task"
)

func TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	_, err = taskClient.TaskCreate(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.TaskCommonResponse{}, nil
}

func SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	resp, err = taskClient.SayHello(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
