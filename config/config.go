package config

import "github.com/spf13/viper"

type (
	Config struct {
		Udp            `mapstructure:"udp"`
		Multithreading `mapstructure:"multithreading"`
		File           `mapstructure:"file"`
	}

	Udp struct {
		Address string `mapstructure:"address"`
	}

	Multithreading struct {
		NumGoroutines int `mapstructure:"num_goroutines"`
		BufferSize    int `mapstructure:"buffer_size"`
	}

	File struct {
		Path string `mapstructure:"path"`
	}
)

var AppSettings Config

func GetConfigs() error {
	if err := initConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&AppSettings)
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
