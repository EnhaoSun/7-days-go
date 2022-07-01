package main

import (
	"fmt"
	"goecache/lru"
)

type String string

func (d String) Len() int {
	return len(d)
}

func main() {
	lru := lru.New(int64(10), func(s string, v lru.Value) {
		fmt.Printf("Evict kv[%s, %v]\n", s, v)
	})
	lru.Add("key1", String("1234"))
	fmt.Println(lru.Get("key1"))
	lru.Add("key2", String("2345"))
	fmt.Println(lru.Get("key2"))
}
