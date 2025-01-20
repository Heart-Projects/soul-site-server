package factory

import (
	"com.sj/admin/pkg/options"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"sync"
	"time"
)

var one sync.Once

// InitialMysqlFactory  初始化数据库工厂
func InitialMysqlFactory(options *options.MysqlOptions) error {
	var error error
	one.Do(func() {
		// 初始化数据库
		dsn := options.Dsn
		if "" == dsn {
			logrus.Error("读取数据库配置信息失败")
			error = fmt.Errorf("can't read database config")
		}

		// 配置日志
		dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             1 * time.Second,
				Colorful:                  true,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Info,
			})

		// 打开数据库
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:               dsn,
			DefaultStringSize: 256,
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			Logger:         dbLogger,
		})

		if err != nil {
			logrus.Error("数据库初始化出错")
			error = err
		}

		// 设置数据库连接池
		dbConnection, err := db.DB()
		// 设置最小空闲数
		dbConnection.SetMaxIdleConns(10)
		// 设置最大打开的连接数
		dbConnection.SetMaxOpenConns(100)
		// 设置连接最大空闲时间
		dbConnection.SetConnMaxLifetime(time.Minute * 30)
		// 导出
		ok, err := StoreInstanceWithType[*gorm.DB](db)
		logrus.Info("初始化数据库工厂成功")
		if ok {
			error = nil
		} else {
			logrus.Error("初始化数据库工厂出错", err)
		}
	})

	return error
}

func GetDbInstance() *gorm.DB {
	return GetInstanceWithType[*gorm.DB]()
}
