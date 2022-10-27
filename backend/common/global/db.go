package global

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// Mysql 配置MySQl数据库

func InitMysql() *gorm.DB {
	m := Cfg.Mysql
	var dsn = fmt.Sprintf("%s:%s@%s", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return nil
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}
	//GORM 使用 database/sql 维护连接池
	sqlDb.SetMaxIdleConns(10)           //设置空闲连接池中连接的最大数量
	sqlDb.SetMaxOpenConns(100)          //SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	return db
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Cfg.Redis.Url,
		Password: Cfg.Redis.Password, // no password set
		DB:       Cfg.Redis.DB,       // use default DB
	})
	return rdb
}
