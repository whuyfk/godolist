package service

import (
	"context"
	"github.com/whuyfk/godolist/app/task/internal/model"
	"github.com/whuyfk/godolist/app/task/internal/query"
	"gorm.io/gorm"
	"sync"
	"time"

	pb "github.com/whuyfk/godolist/idl/pb/task"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
	pb.UnimplementedTaskServiceServer
	db *gorm.DB
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{db: query.InitDb()}
	})
	return TaskSrvIns
}

func (taskSrv *TaskSrv) SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloResponse, err error) {
	resp = &pb.HelloResponse{}
	resp.Msg = "hello"
	return
}

func (taskSrv *TaskSrv) TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskCommonResponse, err error) {
	task := &model.Task{
		TaskID:    req.TaskId,
		UserID:    req.UserId,
		Title:     req.Title,
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	q := query.Use(taskSrv.db)
	err = q.Task.WithContext(ctx).Create(task)
	if err != nil {
		return nil, err
	}
	return &pb.TaskCommonResponse{}, nil
}
