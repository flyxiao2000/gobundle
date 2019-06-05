package main

import (
	"flag"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	addr, hkey, hfield, keys, flushall, hlen string
	status, hdel, hset, get, dbsize          bool
)

func init() {
	flag.StringVar(&addr, "addr", "127.0.0.1:6379", "redis address")
	flag.StringVar(&hkey, "hkey", "", "hashmap key name")
	flag.StringVar(&hfield, "hfield", "", "hashmap field name")
	flag.StringVar(&keys, "keys", "", "hashmap field name")
	flag.StringVar(&hlen, "hlen", "", "hashmap field len")

	flag.StringVar(&flushall, "flushall", "", "redis flushall, ok")

	flag.BoolVar(&dbsize, "dbsize", false, "hsize from redis")

	flag.BoolVar(&hset, "hset", false, "hset value to redis")
	flag.BoolVar(&get, "get", false, "get key value from redis")

	flag.BoolVar(&status, "status", false, "show redis info and config")
	flag.BoolVar(&hdel, "hdel", false, "delet hashmap")
}

func redisdbsize(client *redis.Client) {
	fmt.Println("dbsize ", dbsize)

	if dbsize {
		if val, err := client.DbSize().Result(); err != nil {
			fmt.Println("redis dbsize ", err)
		} else {
			fmt.Println("***************************************")
			fmt.Println(val)
			fmt.Println("***************************************")
		}
	}
}

func redishset(client *redis.Client) {
	fmt.Println("hset ", hset)
	if hset {
		if hkey != "" && hfield != "" {
			if val, err := client.HSet(hkey, hfield, "I am hello").Result(); err != nil {
				fmt.Println("redis hset ", err)
			} else {
				fmt.Println("***************************************")
				fmt.Println(val)
				fmt.Println("***************************************")
			}
		} else {
			fmt.Println("must input hkey and hfield")
		}
	}
}

func redishlen(client *redis.Client) {
	fmt.Println("hlen ", hlen)

	if hlen != "" {
		if val, err := client.HLen(hlen).Result(); err != nil {
			fmt.Println("redis hlen ", err)
		} else {
			fmt.Println("***************************************")
			fmt.Println(val)
			fmt.Println("***************************************")
		}
	}
}

func redisflushall(client *redis.Client) {
	fmt.Println("flushall ", flushall)

	if flushall == "ok" {
		client.FlushAll()
	}
}

func redisdel(client *redis.Client) {
	fmt.Println("hdel ", hdel)

	if hdel {
		client.HDel(hkey, hfield)
	}
}

func redisstatus(client *redis.Client) {
	fmt.Println("status ", status)

	if status {
		if val, err := client.Info().Result(); err != nil {
			fmt.Println("redis info ", err)
		} else {
			fmt.Println("***************************************")
			fmt.Println(val)
			fmt.Println("***************************************")
		}
		if val, err := client.ConfigGet("*").Result(); err != nil {
			fmt.Println("redis config ", err)
		} else {
			fmt.Println("***************************************")
			fmt.Println(val)
			fmt.Println("***************************************")
		}
	}
}

func rediskeys(client *redis.Client) {
	fmt.Println("keys ", keys)

	if keys != "" {
		fmt.Println(keys)
		if val, err := client.Keys(keys).Result(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("***************************************")
			fmt.Println(val)
			fmt.Println("***************************************")
		}
	}
}

func redishashmap(client *redis.Client) {
	fmt.Println("get ", get)
	fmt.Println("hkey ", hkey)
	fmt.Println("hfield ", hfield)

	if get {
		if hkey != "" && hfield != "" {
			if val, err := client.HGet(hkey, hfield).Result(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("***************************************")
				fmt.Println(val)
				fmt.Println("***************************************")
			}
		} else if hkey != "" && hfield == "" {
			if val, err := client.HGetAll(hkey).Result(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("***************************************")
				fmt.Println(val)
				fmt.Println("***************************************")
			}
		}
	}
}

func main() {
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	redisstatus(client)
	rediskeys(client)
	redishashmap(client)
	redisflushall(client)
	redishlen(client)
	redishset(client)
	redisdbsize(client)
}
