package rpc

import (
	"github.com/whuyfk/godolist/idl/pb/task"
	"google.golang.org/grpc"
	"sync"
)

var (
	taskClient task.TaskServiceClient
	once       sync.Once
)

func InitRpc() {
	once.Do(func() {
		conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
		taskClient = task.NewTaskServiceClient(conn)
	})
}
