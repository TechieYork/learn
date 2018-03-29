package main

import (
	"fmt"
	"context"
)

func main() {
	fmt.Println("====== Test context ======")

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
