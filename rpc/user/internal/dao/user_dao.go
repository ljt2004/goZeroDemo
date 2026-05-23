package dao

import (
	"user-grpc/internal/model"

	"gorm.io/gorm"
)

// UserDao 数据库操作对象
type UserDao struct {
	db *gorm.DB
}

// NewUserDao 创建DAO
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// 根据ID查用户 —— 所有数据库操作都写这里
func (d *UserDao) GetUserById(id int64) (*model.User, error) {
	var user model.User
	err := d.db.Select(
		model.User{}.Username,
		model.User{}.NickName,
		model.User{}.Phone,
		model.User{}.Email,
	).Where("id = ?", id).First(&user).Error
	return &user, err
}

// 根据用户名查用户
func (d *UserDao) GetUserByUsername(userId int64) (*model.User, error) {
	var user model.User
	err := d.db.Select(
		model.User{}.Username,
		model.User{}.NickName,
		model.User{}.Phone,
		model.User{}.Email,
	).Where("id = ?", userId).First(&user).Error
	return &user, err
}

// 根据手机号查用户
func (d *UserDao) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	err := d.db.Select("phone").Where("phone = ?", phone).First(&user).Error
	return &user, err
}

// 用户注册
func (d *UserDao) CreateUser(user *model.User) error {
	return d.db.Create(&user).Error
}
