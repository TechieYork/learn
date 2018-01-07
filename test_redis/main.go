package main

import (
    "fmt"

    goredis "github.com/go-redis/redis"
    redigo "github.com/garyburd/redigo/redis"
)

func testGoRedisKeys() {
    fmt.Println("====== Test go-redis keys ======")

    //Connect
    client := goredis.NewClient(&goredis.Options{
        Addr:     "172.16.101.128:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    defer client.Close()

    pong, err := client.Ping().Result()

    if err != nil {
    	fmt.Println("goredis ping error:", err)
    	return
	} else {
		fmt.Println("ping success! pong:", pong)
	}

	//Test set
	err = client.Set("test-go-redis-key", "test set!", 0).Err()

	if err != nil {
		fmt.Println("set failed! error:", err)
		return
	} else {
		fmt.Println("set success!")
	}

	//Test append
	ret, err := client.Append("test-go-redis-key", "+").Result()

	if err != nil {
		fmt.Println("append failed! error:", err)
		return
	} else {
		fmt.Println("append success:", ret)
	}

	//Test get
	data, err := client.Get("test-go-redis-key").Result()

	if err != nil {
		fmt.Println("get failed! error:", err)
		return
	} else {
		fmt.Println("get success:", data)
	}

	/**********************
	 * There are lots of other commands
	 * Check "https://redis.io/commands" to get further information
	 **********************/
}

func testGoRedisHash() {
	fmt.Println("====== Test go-redis hash ======")

	//Connect
	client := goredis.NewClient(&goredis.Options{
		Addr:     "172.16.101.128:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("goredis ping error:", err)
		return
	} else {
		fmt.Println("ping success! pong:", pong)
	}

	//Test hash set
	ok, err := client.HSet("test-go-redis-hash", "name", "techie").Result()

	if err != nil {
		fmt.Println("hash set error:", err)
		return
	} else {
		fmt.Println("hash set success! ok:", ok)
	}

	//Test hash get
	name, err := client.HGet("test-go-redis-hash", "name").Result()

	if err != nil {
		fmt.Println("hash get error:", err)
		return
	} else {
		fmt.Println("hash get success! name:", name)
	}

	//Test hash get - not exist
	name, err = client.HGet("test-go-redis-hash-not-exist", "name").Result()

	if err != nil {
		fmt.Println("hash get error:", err)
	} else {
		fmt.Println("hash get success! name:", name)
	}

	//Test hash del
	count, err := client.HDel("test-go-redis-hash", "name").Result()

	if err != nil {
		fmt.Println("hash del error:", err)
		return
	} else {
		fmt.Println("hash del success! count:", count)
	}
}

func testGoRedisSet() {
	fmt.Println("====== Test go-redis set ======")

	//Connect
	client := goredis.NewClient(&goredis.Options{
		Addr:     "172.16.101.128:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("goredis ping error:", err)
		return
	} else {
		fmt.Println("ping success! pong:", pong)
	}

	//Test set add
	count, err := client.SAdd("test-go-redis-set", "york", "techie").Result()

	if err != nil {
		fmt.Println("set add error:", err)
		return
	} else {
		fmt.Println("set add success! count:", count)
	}

	//Test set is member
	ok, err := client.SIsMember("test-go-redis-set", "york").Result()

	if err != nil {
		fmt.Println("set is member error:", err)
		return
	} else {
		fmt.Println("set is member success! ok:", ok)
	}

	//Test set is member - not exist
	ok, err = client.SIsMember("test-go-redis-set", "haha").Result()

	if err != nil {
		fmt.Println("set is member error:", err)
	} else {
		fmt.Println("set is member success! ok:", ok)
	}

	//Test set members
	members, err := client.SMembers("test-go-redis-set").Result()

	if err != nil {
		fmt.Println("set members error:", err)
	} else {
		fmt.Println("set members success! members:", members)
	}

	//Test set remove
	count, err = client.SRem("test-go-redis-set", "york", "techie").Result()

	if err != nil {
		fmt.Println("set remove error:", err)
		return
	} else {
		fmt.Println("set remove success! count:", count)
	}
}

func testGoRedisZset() {
	fmt.Println("====== Test go-redis zset ======")

	//Connect
	client := goredis.NewClient(&goredis.Options{
		Addr:     "172.16.101.128:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("goredis ping error:", err)
		return
	} else {
		fmt.Println("ping success! pong:", pong)
	}

	//Test zset add
	count, err := client.ZAdd("test-go-redis-zset",
		goredis.Z{Score:99, Member:"techie"},
		goredis.Z{Score:98, Member:"york"},
		goredis.Z{Score:97, Member:"jenny"},
		goredis.Z{Score:96, Member:"monkey"},
		goredis.Z{Score:95, Member:"rabbit"}).Result()

	if err != nil {
		fmt.Println("zset add error:", err)
		return
	} else {
		fmt.Println("zset add success! count:", count)
	}

	//Test zset rev range by score
	ok, err := client.ZRevRangeByScore("test-go-redis-zset",
		goredis.ZRangeBy{Max:"+inf", Min:"-inf", Count:3}).Result()

	if err != nil {
		fmt.Println("zset rev range by score error:", err)
		return
	} else {
		fmt.Println("zset rev range by score success! ok:", ok)
	}

	//Test zset rev rank
	rank, err := client.ZRevRank("test-go-redis-zset", "jenny").Result()

	if err != nil {
		fmt.Println("zset rev rank error:", err)
		return
	} else {
		fmt.Println("zset rev rank success! rank:", rank + 1)
	}

	//Test zset score
	score, err := client.ZScore("test-go-redis-zset", "jenny").Result()

	if err != nil {
		fmt.Println("zset score error:", err)
		return
	} else {
		fmt.Println("zset score success! score:", score)
	}
}

func testGoRedis() {
	testGoRedisKeys()
	testGoRedisHash()
	testGoRedisSet()
	testGoRedisZset()
}

func testRedigo() {
    fmt.Println("====== Test redigo ======")

    //Connect
	client, err := redigo.Dial("tcp", "172.16.101.128:6379")

	if err != nil {
		fmt.Println("redigo dial failed! error:", err)
	}

	defer client.Close()

	//Test set
	_, err = client.Do("SET", "test-redigo", "test set!")

	if err != nil {
		fmt.Println("set failed! error:", err)
		return
	} else {
		fmt.Println("set success!")
	}

	//Test append
	reply, err := redigo.Int(client.Do("APPEND", "test-redigo", "+"))

	if err != nil {
		fmt.Println("append failed! error:", err)
		return
	} else {
		fmt.Println("append success:", reply)
	}

	//Test get
	result, err := redigo.String(client.Do("GET", "test-redigo"))

	if err != nil {
		fmt.Println("get failed! error:", err)
		return
	} else {
		fmt.Println("get success:", result)
	}
}

func main() {
	fmt.Println("====== Test redis ======")

	testGoRedis()

	//testRedigo()
}
