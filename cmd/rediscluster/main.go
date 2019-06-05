package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"192.168.150.220:7000",
			"192.168.150.220:7001",
			"192.168.150.220:7002",
			//			"192.168.150.220:7003",
			//			"192.168.150.220:7004",
			//			"192.168.150.220:7005",
		},
	})
	redisdb.HSet("myhkey", "myhfield", "i am hello")
	redisdb.HSet("myhkey", "myhfield2", "i am hello")
	fmt.Println(redisdb.HGet("myhkey", "myhfield").Result())
	fmt.Println(redisdb.Ping().String())

}
