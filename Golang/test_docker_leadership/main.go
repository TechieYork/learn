package main

import (
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/leadership"
	"flag"
	"fmt"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/zookeeper"
)

var (
	nodeName = flag.String("node", "unknown", "Node's name")
	kv = flag.String("kv", string(store.ZK), "kv's type, default is 'zk'(available for consul, etcd, zk)")
	address = flag.String("address", "127.0.0.1:2181", "kv's address, default is '127.0.0.1:2181")
)
func main() {
	flag.Parse()

	// Create a store using pkg/store.
	consul.Register()
	etcd.Register()
	zookeeper.Register()

	client, err := libkv.NewStore(store.Backend(*kv), []string{*address}, &store.Config{})

	if err != nil {
		panic(err)
	}

	me := leadership.NewCandidate(client, "service/test_leadership/leader", *nodeName, 15 * time.Second)
	electedCh, _ := me.RunForElection()

	for isElected := range electedCh {
		// This loop will run every time there is a change in our leadership
		// status.

		if isElected {
			// We won the election - we are now the leader.
			// Let's do leader stuff, for example, sleep for a while.
			fmt.Printf("[%v]I won the election! I'm now the leader\r\n", *nodeName)
			time.Sleep(10 * time.Second)

			// Tired of being a leader? You can resign anytime.
			me.Resign()
			time.Sleep(1 * time.Second)
		} else {
			// We lost the election but are still running for leadership.
			// `elected == false` is the default state and is the first event
			// we'll receive from the channel. After a successful election,
			// this event can get triggered if someone else steals the
			// leadership or if we resign.

			fmt.Printf("[%v]Lost the election, let's try another time\r\n", *nodeName)
		}
	}
}
