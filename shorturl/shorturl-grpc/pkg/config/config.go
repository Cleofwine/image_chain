package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		IP          string `mapstructure:"ip"`
		Port        int    `mapstructure:"port"`
		AccessToken string `mapstructure:"accessToken"`
	} `mapstructure:"server"`
	Redis struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Pwd  string `mapstructure:"pwd"`
	} `mapstructure:"redis"`
	Mysql struct {
		DSN         string `mapstructure:"dsn"`
		MaxLifeTime int    `mapstructure:"max_life_time"`
		MaxOpenConn int    `mapstructure:"max_open_conn"`
		MaxIdleConn int    `mapstructure:"max_idle_conn"`
	} `mapstructure:"mysql"`
	Log struct {
		Level   string
		LogPath string `mapstructure:"logPath"`
	} `mapstructure:"log"`
	ShortDomain     string `mapstructure:"shortDomain"`
	UserShortDomain string `mapstructure:"userShortDomain"`
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
