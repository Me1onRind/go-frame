package task

type Task struct {
	ID           uint64 `gorm:"column:id;primaryKey"`
	TaskName     string
	Args         string
	Status       uint8
	ExecTimes    uint32
	StartTime    uint32
	LastExecTime uint32
	NextExecTime uint32
	ExecInterval uint32
	CTime        uint32 `gorm:"column:ctime"`
	MTime        uint32 `gorm:"column:mtime"`
}
