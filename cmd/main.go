package main

import (
	"com.sj/admin/pkg/app"
	"github.com/sirupsen/logrus"
)

// @title go blog
// @version 1.0
// @description go 编写的博客后端.
// @termsOfService https://aaa.com

// @contact.name shenjin
// @contact.url https://aaa.com
// @contact.email 2434685038@qq.com

// @tag.name api
// @tag.description	api tag
// @tag.docs.url https://aaa.com
// @tag.docs.description This is my blog site

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9990
// @BasePath /

// @schemes http https
// @x-example-key {"key": "value"}

// @description.markdown
func main() {
	logrus.WithField("foo", "bar").Info("hello")
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
	})
	app := app.NewApp("blog", "blog for golang")
	app.Run()
	logrus.Info("Exist app")
}
