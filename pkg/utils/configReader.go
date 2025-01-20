package utils

import (
	"errors"
	"github.com/spf13/viper"
	"sync"
)

var (
	viperInstance *viper.Viper = nil
	once          sync.Once
)

func GetConfig() *viper.Viper {
	if viperInstance == nil {
		once.Do(func() {
			viperInstance = initConfig()
		})
	}
	return viperInstance
}

const envName = "environment"

// initConfig reads in config file and ENV variables if set.
// 读取application.yml和application-{env}.yml文件，并且自动合并
func initConfig() *viper.Viper {
	viper.SetDefault(envName, "dev")

	viper.AddConfigPath("./configs")
	const defaultFileName = "application"
	viper.SetConfigName(defaultFileName)
	viper.SetConfigType("yml")
	defaultFileReadErr := viper.ReadInConfig()
	var configFileNotFoundError *viper.ConfigFileNotFoundError
	if defaultFileReadErr != nil {
		if !errors.As(defaultFileReadErr, &configFileNotFoundError) {
			panic(defaultFileReadErr)
		}
	}
	currentEnv := viper.GetString(envName)
	environmentRelatedFileName := defaultFileName + "-" + currentEnv
	viper.SetConfigName(environmentRelatedFileName)
	err := viper.MergeInConfig()
	if err != nil {
		if !errors.As(err, &configFileNotFoundError) {
			panic(err)
		}
	}
	return viper.GetViper()
}
