package rabbitMq

import (
	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Addr          string        `yaml:"Addr"`
	User          string        `yaml:"User"`
	Password      string        `yaml:"Password"`
	Vhost         string        `yaml:"Vhost"`
	ConnectionMax int           `yaml:"ConnectionMax"`
	ChannelMax    int           `yaml:"ChannelMax"`
	ChannelActive int           `yaml:"ChannelActive"`
	ChannelIdle   int           `yaml:"ChannelIdle"`
	Health        time.Duration `yaml:"Health"`
	Timeout       time.Duration `yaml:"Timeout"`
	Heartbeat     time.Duration `yaml:"Heartbeat"`
}

// 初始化
func GetConfig() *tcr.RabbitSeasoning {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("rabbitMq")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		}
		panic(err)
	}
	var cfg tcr.RabbitSeasoning
	err := v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
