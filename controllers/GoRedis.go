package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
	Redis, singkatan dari Remote Dictionary Server, adalah penyimpanan data nilai utama di dalam memori yang super cepat.
	Umumnya Redis dimanfaatkan sebagai database, cache, manajemen session, message broker, queue, dan
	jenis kebutuhan lainnya yg relevan dengan operasi real-time dan temporary data.
*/

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func RedisMain() {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	rdb := newRedisClient(redisHost, redisPassword)
	fmt.Println("redis client initialized")

	key := "tugas pbp"
	data := "Hello Redis"
	ttl := time.Duration(3) * time.Second

	// store data using SET command
	//atau kita bisa simpen datanya dulu
	//context background = kaya countainer atau tempat keydatattl bersatu
	op1 := rdb.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

	// untuk nge tes code yang 3 detik kita set sebelumnya

	// time.Sleep(time.Duration(4) * time.Second)

	// get data
	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	log.Println("get operation success. result:", res)
}
