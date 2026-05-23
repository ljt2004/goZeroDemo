package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 公共基础模型（所有表都能继承）
type BaseModel struct {
	ID        int64          `gorm:"primaryKey;column:id;comment:雪花ID主键"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;comment:软删除时间"` // gorm 软删除必须用这个
}

// User 用户表
// Bcrypt Salt 盐值，用于密码加密解密
type User struct {
	BaseModel `gorm:"embedded"`

	Username string `gorm:"column:username;type:varchar(32);uniqueIndex:un_username;not null;comment:用户名(登录账号)"`
	NickName string `gorm:"column:nickname;type:varchar(32);not null;default:'-';comment:昵称"`
	Phone    string `gorm:"column:phone;type:varchar(20);uniqueIndex:un_phone;comment:手机号(可登录)"` // 手机号必须 string
	Email    string `gorm:"column:email;type:varchar(64);comment:邮箱"`                            // 邮箱必须 string
	Password string `gorm:"column:password;type:varchar(128);comment:密码(bcrypt加密)"`              // 加密后很长
	Status   int8   `gorm:"column:status;default:1;comment:状态 1:正常 2:白名单 3:黑名单"`
	IsAdmin  bool   `gorm:"column:is_admin;default:false;comment:是否管理员"`
}

// TableName 固定表名
func (User) TableName() string {
	return "user"
}
