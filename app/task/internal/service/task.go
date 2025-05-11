package service

import (
	"context"
	"sync"

	"github.com/whuyfk/godolist/app/task/internal/repository/dao"
	pb "github.com/whuyfk/godolist/idl/pb/task"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
	pb.UnimplementedTaskServiceServer
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (*TaskSrv) TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	resp = new(pb.TaskCommonResponse)
	resp.Code = 200
	dao.NewTaskDao(ctx).Create(req)
	return
}
