package options

import "fmt"

// Options : represent all need config options
type Options struct {
	MysqlOptions  *MysqlOptions
	RedisOptions  *RedisOptions
	ServerOptions *ServerOptions
}

func NewOptions() *Options {
	o := &Options{
		MysqlOptions:  NewMysqlOptions(),
		RedisOptions:  NewRedisOptions(),
		ServerOptions: NewServerOptions(),
	}
	return o
}

func (o *Options) String() string {
	return fmt.Sprintf("mysql options:[%s]  %s redis options:[%s] %s server options:[%s]", o.MysqlOptions, "\n", o.RedisOptions, "\n", o.ServerOptions)
}
