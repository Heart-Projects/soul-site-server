package options

import (
	"com.sj/admin/pkg/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MysqlOptions struct {
	Dsn string
}

func NewMysqlOptions() *MysqlOptions {
	// 初始化数据库
	dsn, err := getDsnFromConfig()
	if err != nil {
		logrus.Error("读取数据库配置信息失败", err)
		return nil
	}
	return &MysqlOptions{
		Dsn: dsn,
	}
}

func getDsnFromConfig() (string, error) {
	var (
		datasourceUsername         = "datasource.username"
		datasourcePassword         = "datasource.password"
		datasourceHost             = "datasource.host"
		datasourcePort             = "datasource.port"
		datasourceDbname           = "datasource.dbname"
		datasourceConnectionConfig = "datasource.connectionConfig"
	)
	username := utils.GetConfig().GetString(datasourceUsername)
	if username == "" {
		return "", fmt.Errorf("数据库用户名为空")
	}
	password := utils.GetConfig().GetString(datasourcePassword)
	host := utils.GetConfig().GetString(datasourceHost)
	port := utils.GetConfig().GetInt(datasourcePort)
	dbname := utils.GetConfig().GetString(datasourceDbname)
	if dbname == "" {
		return "", fmt.Errorf("数据库名为空")
	}
	connectionConfig := utils.GetConfig().GetString(datasourceConnectionConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, dbname, connectionConfig)
	return dsn, nil
}

func (o *MysqlOptions) String() string {
	return fmt.Sprintf(" dsn: %s", o.Dsn)
}
