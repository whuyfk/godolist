package dao

import (
	"fmt"

	"github.com/whuyfk/godolist/app/task/internal/repository/model"
	"github.com/whuyfk/godolist/idl/pb/task"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) TaskCreate(req *task.TaskRequest) {
	t := &model.Task{
		UserId: req.UserId,
		TaskId: req.TaskId,
	}
	if err := dao.Model(model.Task{}).Create(&t).Error; err != nil {
		fmt.Println(err)
	}
}
