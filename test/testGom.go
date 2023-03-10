package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
// 它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
var ctx = context.Background()

var DB *redis.Client

func main() {
	rdb := redis.NewClient(&redis.Options{
		// 需要修改成你的配置，本地无需修改
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接成功")
	// 成功连接将其赋值给全局变量
	DB = rdb
	res1, err := DB.Set(ctx, "name", "hua", time.Minute).Result()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("insert sucess:%s\n", res1)

}

func redisGet(key string) {

}
