package main

import (
	"fmt"
	"context"
	"time"
)

func testContextValue() {
	type DataKey struct {
		Key string
	}

	type DataInfo struct {
		Content string
	}

	c := context.Background()
	c1 := context.WithValue(c, DataKey{"key1"}, DataInfo{"111"})
	c2 := context.WithValue(c1, DataKey{"key2"}, DataInfo{"222"})
	c3 := context.WithValue(c2, DataKey{"key1"}, DataInfo{"333"})

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	infoInterface := c3.Value(DataKey{"key1"})

	if infoInterface != nil {
		info := infoInterface.(DataInfo)
		fmt.Println(info.Content)
	}
}

func testContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	for index := 0; index < 5; index++ {
		go func(ctx context.Context, index int) {
			select {
			case <- ctx.Done():
				fmt.Printf("Index [%v] ctx.Done()\r\n", index)
			case <- time.After(time.Second * 2):
				fmt.Printf("Index [%v] Time after 2 seconds\r\n")
			}
		}(ctx, index)
	}

	cancel()

	time.Sleep(time.Second * 5)
}

func main() {
	fmt.Println("====== Test context ======")

	//testContextValue()
	testContextCancel()
}
