package audio

type Audio struct {
	ID         uint64 `gorm:"column:id;primaryKey"`
	Filename   string `gorm:"column:filename"`
	StoreIndex string `gorm:"store_index"`
	CTime      uint32 `gorm:"column:ctime"`
	MTime      uint32 `gorm:"column:mtime"`
}

func (a *Audio) TableName() string {
	return "audio_tab"
}
