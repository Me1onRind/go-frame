package base

type Model struct {
	ID    uint64 `gorm:"column:id"`
	CTime uint32 `gorm:"column:ctime"`
	MTime uint32 `gorm:"column:mtime"`
}
