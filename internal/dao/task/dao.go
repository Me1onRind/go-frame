package task

import (
	"go-frame/global"
	"go-frame/internal/constant/task_constant"
	"go-frame/internal/core/custom_ctx"
	"go-frame/internal/core/errcode"
	"go-frame/internal/utils/date"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TaskDao struct {
}

func NewTaskDao() *TaskDao {
	t := &TaskDao{}
	return t
}

func (t *TaskDao) Create(ctx *custom_ctx.Context, task *Task) *errcode.Error {
	db := ctx.WriteDB(global.DefaultDB)
	if err := db.Save(task).Error; err != nil {
		ctx.Logger().Error("Create task record failed", zap.Any("task", task), zap.Error(err))
		return errcode.DBError.WithError(err)
	}
	return nil
}

func (t *TaskDao) BatchUpdateExecInfo(ctx *custom_ctx.Context, ids []uint64) (e *errcode.Error) {
	db := ctx.WriteDB(global.DefaultDB)
	now := date.UnixTime()
	db = db.Model(&Task{}).Where("id IN ?", ids).Updates(map[string]interface{}{
		"last_exec_time": now,
		"next_exec_time": gorm.Expr("next_exec_time + exec_internal"),
		"exec_times":     gorm.Expr("exec_times + 1"),
	})
	if err := db.Error; err != nil {
		ctx.Logger().Error("Update task exec_time failed", zap.Any("ids", ids), zap.Error(err))
		return errcode.DBError.WithError(err)
	}

	return nil
}

func (t *TaskDao) GetValidExecTasks(ctx *custom_ctx.Context, limit int) ([]*Task, *errcode.Error) {
	var result []*Task

	now := date.UnixTime()
	db := ctx.ReadDB(global.DefaultDB)
	if err := db.Where("next_exec_time < ?", now).Where("status = ?", task_constant.Doing).Limit(limit).Find(&result).Error; err != nil {
		ctx.Logger().Error("Get valid exec tasks failed", zap.Error(err))
		return nil, errcode.DBError.WithError(err)
	}

	return result, nil
}
