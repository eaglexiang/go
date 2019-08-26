/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-16 16:45:10
 * @LastEditTime: 2019-04-05 12:30:38
 */
package eaglecache

import (
	"testing"
	"time"
)

func Test_cacheNode_TTL(t *testing.T) {
	node := createcacheNode("testKey", time.Second)
	if node.overTTL() {
		t.Error("cache node is not over-ttl now")
	}
	time.Sleep(time.Second)
	if !node.overTTL() {
		t.Error("cache node should be timeout")
	}
}

func Test_cacheNode_Update(t *testing.T) {
	node := createcacheNode("testKey")
	start := time.Now()
	go func() {
		time.Sleep(time.Second * 1)
		node.Update("testValue")
	}()
	value, err := node.Wait4Value()
	if err != nil {
		t.Error(err)
	}
	if time.Since(start) < time.Second*1 {
		t.Error("too short after start(<1s)")
	}
	if value != "testValue" {
		t.Error("error value: ", value)
	}
}
