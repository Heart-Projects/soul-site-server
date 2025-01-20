package app

import (
	"com.sj/admin/pkg/controller"
	"com.sj/admin/pkg/factory"
	"com.sj/admin/pkg/options"
	"com.sj/admin/pkg/server"
	"github.com/sirupsen/logrus"
)

type Application struct {
	name        string
	description string
	options     *options.Options
	apiServer   *server.GenericAPIServer
}

func NewApp(name string, description string) *Application {
	options := options.NewOptions()
	app := &Application{
		name:        name,
		description: description,
		options:     options,
		apiServer:   server.New(options.ServerOptions),
	}
	logrus.Info(app.options)
	return app
}

func (a *Application) PrepareRun() {
	// 初始化数据库
	err := factory.InitialMysqlFactory(a.options.MysqlOptions)
	if err != nil {
		logrus.Error("初始化数据库失败, error is %s", err)
		return
	}
	logrus.Info("开始注册路由")
	controller.InitRouters(a.apiServer.Engine)

}

func (a *Application) Run() {
	a.PrepareRun()
	a.apiServer.Run()
}
