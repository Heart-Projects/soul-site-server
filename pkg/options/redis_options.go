package options

import (
	"com.sj/admin/pkg/utils"
	"fmt"
)

type RedisOptions struct {
	Host     string
	Port     int
	Db       int
	Username string
	Password string
}

func NewRedisOptions() *RedisOptions {
	const (
		redisHost     = "redis.host"
		redisPort     = "redis.port"
		redisUsername = "redis.username"
		redisPassword = "redis.password"
		redisDb       = "redis.db"
	)
	utils.GetConfig().SetDefault(redisHost, "127.0.0.1")
	utils.GetConfig().SetDefault(redisPort, 6379)
	utils.GetConfig().SetDefault(redisUsername, "root")
	utils.GetConfig().SetDefault(redisPassword, "")
	utils.GetConfig().SetDefault(redisDb, 0)
	return &RedisOptions{
		Host:     utils.GetConfig().GetString(redisHost),
		Password: utils.GetConfig().GetString(redisPassword),
		Port:     utils.GetConfig().GetInt(redisPort),
		Db:       utils.GetConfig().GetInt(redisDb),
		Username: utils.GetConfig().GetString(redisUsername),
	}
}

func (o *RedisOptions) String() string {
	return fmt.Sprintf("host: %s, port: %d, username: %s, password: %s, db: %d", o.Host, o.Port, o.Username, o.Password, o.Db)
}
