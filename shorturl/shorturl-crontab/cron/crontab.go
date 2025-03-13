package cron

import "github.com/robfig/cron/v3"

func Run() {
	initalLoad() // 防止当天部署没到凌晨三点就没有数据的尴尬
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 3 * * *", setUrlMapMaxID)
	c.Run()
}

func initalLoad() {
	setUrlMapMaxID()
}
