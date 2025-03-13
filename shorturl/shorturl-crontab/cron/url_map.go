package cron

import (
	"context"
	"fmt"
	"shorturl-crontab/data"
	"shorturl-crontab/pkg/db/mysql"
	"shorturl-crontab/pkg/db/redis"
	"shorturl-crontab/pkg/log"
	"time"
)

const DefaultUrlMapTTL = 86400 * 30

func setUrlMapMaxID() {
	tables := []string{"url_map", "url_map_user"}
	redisPool := redis.GetPool()
	client := redisPool.Get()
	defer redisPool.Put(client)

	d := data.NewData(mysql.GetDB())
	for _, t := range tables {
		id, err := d.GetMaxID(t)
		if err != nil {
			log.Error(err)
			continue
		}
		key := redis.GetKey(fmt.Sprintf("%s_%s", t, "maxid"))
		err = client.SetEx(context.Background(), key, id, time.Second*time.Duration(DefaultUrlMapTTL)).Err()
		if err != nil {
			log.Error(err)
			continue
		}
	}
}
