package task

type RetryTaskInfo struct {
	TaskName   string
	Args       []interface{}
	RetryTimes int
}

type RetryTask interface {
	Create(task *RetryTaskInfo)
}
