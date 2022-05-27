package Config

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisCache *redis.Client

func RedisInit() {
	//RedisClient = redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
	//	Password: viper.GetString("redis.auth"),
	//	DB:       0,
	//})
	//time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	redisCache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
		Password: "",
		DB:       0,
	})

	_, err := redisCache.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}

func SetStr(key, value string, time time.Duration) (err *redis.StatusCmd) {
	err = redisCache.Set(key, value, time)
	if err != nil {
		//beego.Error("set key:", key, ",value:", value, err)
	}
	return
}

func GetStr(key string) (value string) {
	v := redisCache.Get(key)
	value, err := v.Result() //这里的转换很重要，Get返回的是interface
	if err != nil {
		fmt.Println("获取值错误")
		return
	}

	return
}

func DelKey(key string) (err *redis.IntCmd) {
	err = redisCache.Del(key)
	return
}

func CheckIsExist(key string) (err *redis.IntCmd) {
	err = redisCache.Exists(key)
	return
}
