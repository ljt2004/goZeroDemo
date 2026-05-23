package svc

import (
	"fmt"
	"time"
	"user-grpc/internal/config"
	"user-grpc/internal/dao"
	"user-grpc/internal/model"

	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	Snowflake *snowflake.Node // 雪花节点
	UserDao   *dao.UserDao
}

func NewServiceContext(c config.Config) *ServiceContext {

	// mysql
	mysqlConf := c.MySQLConf
	var db *gorm.DB
	var err error
	if mysqlConf.Enable {
		db, err = creteDbClient(mysqlConf)
		if err != nil {
			db = nil
		}
	}

	// 初始化雪花算法
	// 注意：每个微服务机器号不同 1~1023
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		DB:        db,
		Snowflake: node,
		UserDao:   dao.NewUserDao(db),
	}
}

// 初始化数据库
func creteDbClient(mysqlConf config.MySQLConf) (*gorm.DB, error) {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
		mysqlConf.CharSet,
		mysqlConf.ParseTime,
		mysqlConf.TimeZone)
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mysqlConf.Gorm.TablePrefix,   // such as: prefix_tableName
			SingularTable: mysqlConf.Gorm.SingularTable, // such as zero_user, not zero_users
		},
	})
	if err != nil {
		logx.Errorf("create mysql db failed, err: %v", err)
		return nil, err
	}

	// auto sync table structure, no need to create table
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		logx.Errorf("automigrate table failed, err: %v", err)
	}

	logx.Info("init mysql client instance success.")

	sqlDB, err := db.DB()
	if err != nil {
		logx.Errorf("mysql set connection pool failed, err: %v.", err)
		return nil, err
	}
	sqlDB.SetMaxOpenConns(mysqlConf.Gorm.MaxOpenConns)
	sqlDB.SetMaxIdleConns(mysqlConf.Gorm.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.Gorm.ConnMaxLifetime) * time.Second)

	return db, nil
}
