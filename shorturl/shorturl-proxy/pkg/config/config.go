package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Http struct {
		IP   string `mapstructure:"ip"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"http"`
	DependOn struct {
		ShortUrl struct {
			Address     string `mapstructure:"address"`
			AccessToken string `mapstructure:"accessToken"`
		} `mapstructure:"shortUrl"`
	} `mapstructure:"dependOn"`
	Log struct {
		Level   string
		LogPath string `mapstructure:"logPath"`
	} `mapstructure:"log"`
}

var conf *Config

func InitConfig(filePath string, typ ...string) {
	v := viper.New()
	v.SetConfigFile(filePath)
	if len(typ) > 0 {
		v.SetConfigType(typ[0])
	}
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf = &Config{}
	err = v.Unmarshal(conf)
	if err != nil {
		log.Fatal(err)
	}
	// 配置热更新
	v.OnConfigChange(func(in fsnotify.Event) {
		v.Unmarshal(conf)
	})
	v.WatchConfig()
}

func GetConfig() *Config {
	return conf
}
