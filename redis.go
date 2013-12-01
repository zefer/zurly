package main

import (
	redislib "github.com/hoisie/redis"
	"log"
	"os"
)

var redis redislib.Client

type RedisConn struct {
	Client   redislib.Client
	Address  string
	Password string
}

func connectToRedis() {
	conn := RedisConn{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	conn.connect()
	conn.initCounter()
	redis = conn.Client
}

func (this *RedisConn) connect() {
	if this.Address != "" {
		log.Printf("Connecting to redis at %v", this.Address)
		this.Client = redislib.Client{
			Addr:     this.Address,
			Password: this.Password,
		}
	}
}

func (this *RedisConn) initCounter() {
	counter, _ := this.Client.Get("zurl:counter")
	if counter == nil {
		log.Print("Initialising counter")
		this.Client.Set("zurl:counter", []byte("0"))
	} else {
		log.Printf("Counter is " + string(counter))
	}
}
