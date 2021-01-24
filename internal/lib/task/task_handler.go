package task

import (
	"encoding/json"
	"fmt"
	"go-frame/internal/core/errcode"
	"reflect"
)

var (
	taskRunners = map[string]TaskRunner{}
)

type TaskHandler func(args []interface{}) *errcode.Error

type TaskRunner func(args [][]byte) *errcode.Error

func NewTaskRunner(fn TaskHandler, paramTypes ...interface{}) TaskRunner {
	return func(args [][]byte) *errcode.Error {

		if len(args) != len(paramTypes) {
			return nil
		}

		params := make([]interface{}, len(args))
		for k, paramType := range paramTypes {
			pt := reflect.TypeOf(paramType)
			pt = pt.Elem()
			value := reflect.New(pt).Interface()

			if err := json.Unmarshal(args[k], value); err != nil {
				return nil
			}

			params = append(params, value)
		}

		return fn(params)
	}
}

func AddTaskRunner(taskName string, tr TaskRunner) {
	if _, ok := taskRunners[taskName]; ok {
		panic(fmt.Sprintf("Duplicate add task runner, task name:%s", taskName))
	}
	taskRunners[taskName] = tr
}
