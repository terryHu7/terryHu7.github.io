package main

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	/*
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		key := "mykey"
		value := "myvalue"

		setNXResult, err := rdb.SetNX(ctx, key, value, 0).Result()
		if err != nil {
			fmt.Println("SETNX failed:", err)
			return
		}

		if setNXResult {
			fmt.Println("SETNX succeeded: key does not exist")
		} else {
			fmt.Println("SETNX failed: key already exists")
		}
	*/
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println("lock succ, do your business logic")

	time.Sleep(time.Second * 10)

	// do some thing
	l.Unlock()
	println("unlock succ, finish business logic")
}
