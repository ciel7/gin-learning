package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

// initializeDatabase 初始化连接Redis
func initializeRedisClient() (err error) {
	// 方法一
	//rdb = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       0,
	//	PoolSize: 100, // 连接池大小
	//})

	// 方法二
	opt, err := redis.ParseURL("redis://root:@localhost:6379/0")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return
}

func hGetAllDemo() {
	v := rdb.HGetAll(ctx, "user").Val()
	fmt.Println("v = ", v) // map[]

	v1 := rdb.HMGet(ctx, "user", "name", "age").Val()
	fmt.Println("v1 = ", v1) //  [<nil> <nil>]

	v2 := rdb.HGet(ctx, "user", "age").Val()
	fmt.Println("v2 = ", v2) // ""
}

func main() {
	if err := initializeRedisClient(); err != nil {
		fmt.Printf("connect to redis failed err： %v \n", err)
		panic(err)
	}
	fmt.Println("connect to redis success")

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("key = %v \n", val)

	val1, err := rdb.Get(ctx, "key1").Result()
	if err == redis.Nil {
		fmt.Println("key1 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("key1 = %v \n", val1)
	}

	hGetAllDemo()
}
