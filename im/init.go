/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package im

import (
	"time"

	"im_app/pkg/config"
	"im_app/pkg/model"
	"im_app/pkg/mq"
	"im_app/pkg/pool"
	"im_app/pkg/redis"
)

func SetupPool() {
	// 启动mysql连接池
	db := model.ConnectDB()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个连接的超时时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	// 启动redis连接池
	redis.InitClient()
	// 启动协程池
	pool.ConnectPool()
	// 启动mq
	mq.ConnectMQ()
}
