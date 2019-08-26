/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-03 14:28:35
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-08-26 23:38:25
 */

package textcache

import (
	"time"

	eaglecache "github.com/eaglexiang/go/cache"
)

// DefaultTextCacheTTL 默认的TextCache生存时间 2h
const DefaultTextCacheTTL = time.Duration(2) * time.Hour

// TextCache DNS缓存
// 必须使用CreateTextCache进行初始化
type TextCache struct {
	value *eaglecache.Cache
}

// CreateTextCache 创建TextCache的方法
// 通过ttl参数指定缓存生存时间，不指定即使用默认值（2h）
func CreateTextCache(ttl ...time.Duration) *TextCache {
	var value *eaglecache.Cache
	if len(ttl) == 0 {
		value = eaglecache.CreateCache(DefaultTextCacheTTL)
	} else if len(ttl) == 1 {
		value = eaglecache.CreateCache(ttl[0])
	} else {
		panic("createTextCache: too many args")
	}
	cache := TextCache{value: value}
	return &cache
}

// Get 获取指定domain的节点
// 以及该节点之前是否存在
func (cache *TextCache) Get(domain string) (node *CacheNode, loaded bool) {
	var _node *eaglecache.CacheNode
	_node, loaded = cache.value.Get(domain)
	node = &CacheNode{node: _node}
	return
}
