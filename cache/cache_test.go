/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-15 21:22:03
 * @LastEditTime: 2019-08-26 23:44:05
 */
package eaglecache

import (
	"testing"
	"time"
)

func Test_Cache_Update(t *testing.T) {
	cache := CreateCache(time.Second)
	defer cache.Close()
	node, loaded := cache.Get("testKey0")
	if loaded {
		t.Error("exsit loaded node")
	}
	go func() {
		node.Update("testValue0")
	}()
	v, err := node.Wait4Value()
	if err != nil {
		t.Error(err)
	}
	if v != "testValue0" {
		t.Error("value for testKey0 should be testValue0: ", v)
	}
}

func Test_Cache_TTL(t *testing.T) {
	cache := CreateCache(time.Second)
	defer cache.Close()
	node, loaded := cache.Get("testKey")
	if loaded {
		t.Error("testKey shouldn't be loaded")
	}
	time.Sleep(time.Second)
	_, err := node.Wait4Value()
	if err == nil {
		t.Error("timeout should occur")
	}
}

func Test_Cache_Get(t *testing.T) {
	c := CreateCache()
	defer c.Close()
	node, loaded := c.Get("key0")
	if node == nil {
		t.Error("node shouldn't be nil")
	}
	if loaded {
		t.Error("key0 should be created")
	}
	node, loaded = c.Get("key0")
	if node == nil {
		t.Error("node shouldn't be nil")
	}
	if !loaded {
		t.Error("key0 should be loaded")
	}
}
