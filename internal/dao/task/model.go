package task

type Task struct {
	ID           uint64 `gorm:"column:id;primaryKey"`
	TaskName     string `gorm:"column:task_name"`
	Args         string `gorm:"args"`
	Status       uint8  `gorm:"status"`
	ExecTimes    uint32 `gorm:"exec_time"`
	StartTime    uint32 `gorm:"start_time"`
	LastExecTime uint32 `gorm:"last_exec_time"`
	NextExecTime uint32 `gorm:"next_exec_time"`
	ExecInterval uint32 `gorm:"exec_interval"`
	CTime        uint32 `gorm:"column:ctime"`
	MTime        uint32 `gorm:"column:mtime"`
}

func (t *Task) TableName() string {
	return "task_tab"
}
