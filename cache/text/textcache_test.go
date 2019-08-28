/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-04-02 23:08:32
 * @LastEditTime: 2019-04-02 23:10:36
 */

package textcache

import "testing"

func Test_TextCache_Get(t *testing.T) {
	c := CreateTextCache()
	node, loaded := c.Get("testKey")
	if node == nil {
		t.Error("node shouldn't be nil")
	}
	if loaded {
		t.Error("testKey shouldn't be loaded")
	}
	node, loaded = c.Get("testKey")
	if node == nil {
		t.Error("node shouldn't be nil")
	}
	if !loaded {
		t.Error("testKey should be loaded")
	}
}
