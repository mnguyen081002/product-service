package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	ConfigDefaultFile = "config/config.yaml"
	configType        = "yaml"
)

type (
	Config struct {
		Debug          bool     `mapstructure:"debug"`
		ContextTimeout int      `mapstructure:"contextTimeout"`
		Server         Server   `mapstructure:"server"`
		Services       Services `mapstructure:"services"`
		Database       Database `mapstructure:"database"`
		Logger         Logger   `mapstructure:"logger"`
		Jwt            Jwt      `mapstructure:"jwt"`
		Kafka          Kafka    `mapstructure:"kafka"`
	}

	Server struct {
		Host     string `mapstructure:"host"`
		Env      string `mapstructure:"env"`
		UseRedis bool   `mapstructure:"useRedis"`
		Port     int    `mapstructure:"port"`
	}

	Database struct {
		Driver   string   `mapstructure:"driver"`
		Postgres Postgres `mapstructure:"postgres"`
		Mongo    Mongo    `mapstructure:"mongo"`
	}

	Postgres struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
		TimeZone string `mapstructure:"timeZone"`
	}

	Mongo struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
		TimeZone string `mapstructure:"timeZone"`
	}

	Jwt struct {
		Secret                string `mapstructure:"secret"`
		AccessTokenExpiresIn  int64  `mapstructure:"accessTokenExpiresIn"`
		RefreshTokenExpiresIn int64  `mapstructure:"refreshTokenExpiresIn"`
		Kid                   string `mapstructure:"kid"`
	}

	Logger struct {
		Enabled bool   `mapstructure:"enabled"`
		Level   string `mapstructure:"level"`
		Format  string `mapstructure:"format"`
		Prefix  string `mapstructure:"prefix"`
	}

	Services struct {
	}

	Kafka struct {
		Host   string `mapstructure:"host"`
		Port   int    `mapstructure:"port"`
		Enable bool   `mapstructure:"enable"`
	}
)

func NewConfig(envPath string) func() *Config {
	return func() *Config {
		initConfig(envPath)
		conf := &Config{}
		err := viper.Unmarshal(conf)
		if err != nil {
			fmt.Printf("unable decode into config struct, %v", err)
		}
		return conf
	}
}

func initConfig(envPath string) {
	viper.SetConfigType(configType)
	pwd, _ := os.Getwd()
	viper.SetConfigFile(pwd + envPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
