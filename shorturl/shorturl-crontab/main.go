package main

import (
	"flag"
	"shorturl-crontab/cron"
	"shorturl-crontab/pkg/config"
	"shorturl-crontab/pkg/db/mysql"
	"shorturl-crontab/pkg/db/redis"
	"shorturl-crontab/pkg/log"
)

var (
	configFile = flag.String("config", "dev.config.yaml", "")
)

func main() {
	flag.Parse()
	config.InitConfig(*configFile)
	cnf := config.GetConfig()
	mysql.InitMysql(cnf)
	redis.InitRedisPool(cnf)
	log.SetLevel(cnf.Log.Level)
	log.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	log.SetPrintCaller(true)
	cron.Run()
}
