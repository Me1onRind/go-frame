package auth

type Auth struct {
	ID        uint64      `gorm:"column:id;primaryKey"`
	AppName   string      `gorm:"column:app_name"`
	AppKey    string      `gorm:"column:app_key"`
	AppSecret string      `gorm:"column:app_secret"`
	ConfigID  uint64      `gorm:"config_id"`
	Config    *AuthConfig `gorm:"foreignkey:id;references:config_id"`
	CTime     uint32      `gorm:"column:ctime"`
	MTime     uint32      `gorm:"column:mtime"`
}

func (a *Auth) TableName() string {
	return "auth"
}

type AuthConfig struct {
	ID         uint64 `gorm:"column:id;primaryKey"`
	ConfigName string `gorm:"column:config_name"`
	Expires    uint32 `gorm:"column:expires"`
	Flag       uint64 `gorm:"column:flag"`
	CTime      uint32 `gorm:"column:ctime"`
	MTime      uint32 `gorm:"column:mtime"`
}

func (a *AuthConfig) TableName() string {
	return "auth_config_tab"
}

type Role struct {
	ID       uint64 `gorm:"column:id;primaryKey" json:"id"`
	RoleID   uint64 `gorm:"column:role_id" json:"role_id"`
	RoleName string `gorm:"column:role_name" json:"role_name"`
	CTime    uint32 `gorm:"column:ctime" json:"ctime"`
	MTime    uint32 `gorm:"column:mtime" json:"mtime"`
}

func (r *Role) TableName() string {
	return "role_tab"
}
