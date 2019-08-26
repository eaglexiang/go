/*
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-21 16:09:37
 * @LastEditTime: 2019-08-26 21:25:33
 */

package eaglecache

import (
	"errors"
	"sync"
	"time"
)

// CacheNode 缓存节点
// CacheNode 必须通过函数createcacheNode进行创建
// 本结构体线程安全
type CacheNode struct {
	sync.Mutex
	key           interface{}
	value         interface{}
	waiting       chan interface{}  // 标识当前Node是否处于阻塞态
	destroyed     chan interface{}  // 标识当前Node是否被销毁
	destroyCaller func(interface{}) // 销毁后的回调，用来帮助外部结构管理索引
	updated       time.Time         // 创建或更新的时间，用来计算超时
	ttl           time.Duration     // 最大TTL
}

// createcacheNode 创建cacheNode
func createcacheNode(key interface{}, ttl ...time.Duration) *CacheNode {
	node := CacheNode{
		key:       key,
		waiting:   make(chan interface{}), // 节点天生阻塞，需要Update之后解除阻塞
		destroyed: make(chan interface{}),
		updated:   time.Now(),
	}
	if len(ttl) == 1 {
		node.ttl = ttl[0]
	} else if len(ttl) > 1 {
		panic("CreatecacheNode: too many args")
	}
	return &node
}

// overTTL 检查TTL是否过期, >0 未过期
func (cache *CacheNode) overTTL() bool {
	if cache.ttl == 0 {
		return false
	}
	d := time.Since(cache.updated)
	return d > cache.ttl
}

// Wait4Value 检查Cache，阻塞于waiting状态，并于状态改变后返回value
func (cache *CacheNode) Wait4Value() (value interface{}, err error) {
	var (
		destroyed = cache.destroyed
		waiting   = cache.waiting
	)
	if destroyed == nil {
		err = errors.New("cache node destroyed")
	}
	if waiting == nil { // 已经解除阻塞
		value = cache.value
		return
	}
	select {
	case <-waiting:
		value = cache.value
	case <-destroyed:
		err = errors.New("cache node destroyed")
	}
	return
}

// Update 更新Cache的Value
// 更新操作会解除节点的阻塞
func (cache *CacheNode) Update(value interface{}) {
	cache.value = value
	cache.updated = time.Now()

	if cache.waiting == nil {
		return
	}
	cache.Lock()
	cache.Unlock()
	if cache.waiting == nil {
		return
	}
	close(cache.waiting)
	cache.waiting = nil
}

// Destroy 销毁Cache
func (cache *CacheNode) Destroy() {
	if cache.destroyed == nil {
		return
	}
	cache.Lock()
	defer cache.Unlock()
	if cache.destroyed == nil {
		return
	}
	close(cache.destroyed)
	cache.destroyed = nil
	if cache.destroyCaller != nil {
		cache.destroyCaller(cache.key)
	}
}
