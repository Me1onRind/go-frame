package task

import (
	"go-frame/internal/constant/task_constant"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"go-frame/internal/dao/task"
)

type RetryTaskInfo struct {
	TaskName     string
	Args         []interface{}
	RetryTimes   uint32
	StartTime    uint32
	ExecInterval uint32
}

func NewRetryTaskService() *RetryTaskService {
	r := &RetryTaskService{
		taskDao: task.NewTaskDao(),
	}
	return r
}

type RetryTaskService struct {
	taskDao *task.TaskDao
}

func (r *RetryTaskService) Create(ctx *custom_ctx.Context, taskInfo *RetryTaskInfo) *errcode.Error {
	task := &task.Task{
		TaskName:     taskInfo.TaskName,
		Args:         "",
		Status:       task_constant.Doing,
		StartTime:    taskInfo.StartTime,
		NextExecTime: taskInfo.StartTime,
		ExecInterval: taskInfo.ExecInterval,
	}
	return r.taskDao.Create(ctx, task)
}
