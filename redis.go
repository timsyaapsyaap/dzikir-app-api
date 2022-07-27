package main

// import (
// 	"crypto/tls"
// 	"fmt"

// 	"github.com/gomodule/redigo/redis"
// )

// var pool = newPool()

// func main() {
// 	client := pool.Get()
// 	defer client.Close()

// 	_, err := client.Do("SET", "mykey", "Hello from redigo!")
// 	if err != nil {
// 		panic(err)
// 	}

// 	value, err := client.Do("GET", "mykey")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%s\n", value)
// }

// func newPool() *redis.Pool {
// 	return &redis.Pool{
// 		MaxIdle:   80,
// 		MaxActive: 12000,
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.DialURL("redis://:pc1d95ff0d910977a783680a8c52015bf4e236ecef94e5ae97d3c4b0abda5bd3a@ec2-54-163-243-115.compute-1.amazonaws.com:9939", redis.DialTLSConfig(&tls.Config{InsecureSkipVerify: true}))
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			return c, err
// 		},
// 	}
// }
