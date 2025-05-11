package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whuyfk/godolist/app/gateway/rpc"
	pb "github.com/whuyfk/godolist/idl/pb/task"
)

func TaskCreate(ctx *gin.Context) {
	req := &pb.TaskRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		fmt.Println(err)
		return
	}
	response, err := rpc.TaskCreate(ctx, req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  nil,
		"data": response,
	})
}

func SayHello(ctx *gin.Context) {
	req := &pb.HelloRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		fmt.Println(err)
		return
	}
	response, err := rpc.SayHello(ctx, req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": response,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  nil,
		"data": response,
	})
}
