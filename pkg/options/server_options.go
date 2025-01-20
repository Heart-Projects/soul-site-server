package options

import (
	"com.sj/admin/pkg/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type ServerOptions struct {
	// Port 启动端口
	Port int
	// Mode gin 启动模式
	Mode string
	// Middlewares 中间件
	Middlewares []string
	// ShutdownTimeout 优雅关闭超时时间
	ShutdownTimeout time.Duration
	// 是否启用探活
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

func NewServerOptions() *ServerOptions {
	const (
		ServerMode            = "server.mode"
		ServerPort            = "server.port"
		ServerHealthz         = "server.healthz"
		ServerEnableProfiling = "server.enableProfiling"
		ServerEnableMetrics   = "server.enableMetrics"
		ServerShutdownTimeout = "server.shutdownTimeout"
	)
	// 配置默认值
	utils.GetConfig().SetDefault(ServerMode, "debug")
	utils.GetConfig().SetDefault(ServerPort, 8099)
	utils.GetConfig().SetDefault(ServerHealthz, true)
	utils.GetConfig().SetDefault(ServerEnableProfiling, true)
	utils.GetConfig().SetDefault(ServerEnableMetrics, true)
	utils.GetConfig().SetDefault(ServerShutdownTimeout, 5*time.Second)
	logrus.Info("port is ", utils.GetConfig().GetInt(ServerPort))
	return &ServerOptions{
		Port:            utils.GetConfig().GetInt(ServerPort),
		Mode:            utils.GetConfig().GetString(ServerMode),
		Middlewares:     []string{},
		ShutdownTimeout: utils.GetConfig().GetDuration(ServerShutdownTimeout) * time.Second,
		Healthz:         utils.GetConfig().GetBool(ServerHealthz),
		EnableProfiling: utils.GetConfig().GetBool(ServerEnableProfiling),
		EnableMetrics:   utils.GetConfig().GetBool(ServerEnableMetrics),
	}

}

func (o *ServerOptions) String() string {
	return fmt.Sprintf(" port: %d, mode: %s", o.Port, o.Mode)
}
