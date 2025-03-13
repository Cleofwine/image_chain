package mysql

import (
	"database/sql"
	"shorturl-crontab/pkg/config"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql(cnf *config.Config) {
	var err error
	if cnf.Mysql.DSN == "" {
		panic("数据库连接字符串为空")
	}
	db, err = sql.Open("mysql", cnf.Mysql.DSN)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Duration(cnf.Mysql.MaxLifeTime) * time.Second)
	db.SetMaxOpenConns(cnf.Mysql.MaxOpenConn)
	db.SetMaxIdleConns(cnf.Mysql.MaxIdleConn)
}

func GetDB() *sql.DB {
	return db
}
