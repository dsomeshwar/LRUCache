package lrucache

import (
	"testing"
	"fmt"
)

func TestLRUCache1(t *testing.T){
	_,err := NewLRUCache(0)
	if err == nil {
		t.Error("Wrong")
	}
}

func TestLRUCache2(t *testing.T){
	lruCache,err := NewLRUCache(2)
	if err != nil {
		t.Error("Wrong")
	}
	lruCache.Set("1","a")
	fmt.Printf("%+v\n",lruCache)
	lruCache.Set("2","b")
	fmt.Printf("%+v\n",lruCache)
	x := lruCache.Get("1")
	lruCache.Set("3","c")
	fmt.Printf("%+v\n",lruCache)

	fmt.Println("x",x)
}

