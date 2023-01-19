package log

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name             string `yaml:"Name"`
	Development      bool   `yaml:"Development"`
	Level            string `yaml:"Level"`
	Format           string `yaml:"Format"`
	Stacktrace       bool   `yaml:"Stacktrace"`
	LinkName         string `yaml:"LinkName"`
	Prefix           string `yaml:"Prefix"`
	Director         string `yaml:"Director"`
	LogRollingPolicy string `yaml:"LogRollingPolicy"`
	LoggerInfoFile   string `yaml:"LoggerInfoFile"`
	LoggerWarnFile   string `yaml:"LoggerWarnFile"`
	LoggerErrorFile  string `yaml:"LoggerErrorFile"`
	MaxAge           int64  `yaml:"MaxAge"`
}

// 初始化
func Init() Logger {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("logger")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		}
		panic(err)
	}
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	logger = newZapLogger(&cfg)
	return logger
}
