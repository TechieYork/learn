package main

import (
	"fmt"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/zookeeper"
)

func main() {
	zookeeper.Register()

	kv, err := libkv.NewStore(
		store.ZK,
		[]string{"172.16.101.128:2181"},
		&store.Config{
			ConnectionTimeout: 10*time.Second,
			},
	)

	if err != nil {
		panic(err)
	}

	key := "test/haha"
	value := "HelloWorld"

	err = kv.Put(key, []byte(value), nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Zookeeper Put, pair.key:%v, pair.value:%v\r\n", key, value)

	pair, err := kv.Get(key)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Zookeeper Get, pair.key:%v, pair.value:%v\r\n", pair.Key, string(pair.Value))
}