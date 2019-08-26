/*
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-21 16:09:17
 * @LastEditTime: 2019-08-26 21:24:55
 */

package eaglecache

import (
	"sync"
	"time"
)

// Cache 缓存
// 必须使用CreateCache函数进行初始化
// 该结构体线程安全
type Cache struct {
	sync.Mutex
	data       sync.Map         // 存放cacheNode的容器
	defaultTTL time.Duration    // 每个新cacheNode的默认ttl，0代表不限制
	cacheNodes sync.Pool        // node对象池，用来节省创建新node时的开销
	destroyed  chan interface{} // 标记Cache是否已摧毁，用来控制定时器的退出
}

// CreateCache 创建Cache
func CreateCache(defaultTTL ...time.Duration) *Cache {
	var ttl time.Duration
	if len(defaultTTL) == 1 {
		ttl = defaultTTL[0]
	} else if len(defaultTTL) > 1 {
		panic("CreateCache: too many args")
	}
	c := &Cache{
		defaultTTL: ttl,
		data:       sync.Map{},
		cacheNodes: sync.Pool{
			New: func() interface{} {
				return createcacheNode(nil, defaultTTL...)
			},
		},
	}

	go c.checkTTL()

	return c
}

// add 增加key,value键值对
// 如果key不存在，则新建节点，并将其设为阻塞态，然后返回存入的节点和false
// 如果key存在，则返回存在的节点和true
func (cache *Cache) tryAdd(key interface{}) (actual *CacheNode, loaded bool) {
	node := cache.getCacheNode(key)
	_actual, loaded := cache.data.LoadOrStore(key, node)
	actual = _actual.(*CacheNode)
	if loaded {
		cache.putCacheNode(node)
	}
	return
}

// Get 拿到key对应的节点
// 并且返回节点是否之前存在
func (cache *Cache) Get(key interface{}) (node *CacheNode, loaded bool) {
	return cache.tryAdd(key)
}

// 检查key是否存在，仅供测试使用
func (cache *Cache) exsit(key interface{}) bool {
	_, loaded := cache.data.Load(key)
	return loaded
}

func (cache *Cache) getCacheNode(key interface{}) (node *CacheNode) {
	node = cache.cacheNodes.Get().(*CacheNode)
	node.key = key
	node.updated = time.Now()
	node.destroyCaller = func(theKey interface{}) {
		cache.data.Delete(theKey)
	}
	return
}

func (cache *Cache) putCacheNode(node *CacheNode) {
	cache.cacheNodes.Put(node)
}

func (cache *Cache) checkTTL() {
	if cache.defaultTTL < time.Second {
		// 间隔过短的检查性能开销过大
		return
	}
	tick := cache.defaultTTL / 3

	var destroyed chan interface{}
	cache.Lock()
	if cache.destroyed == nil {
		cache.destroyed = make(chan interface{})
	}
	destroyed = cache.destroyed
	cache.Unlock()

	for {
		time.Sleep(tick)
		select {
		case <-destroyed:
			break
		default:
			cache.checkEachTTL()
		}
	}
}

func (cache *Cache) checkEachTTL() {
	cache.data.Range(func(_, v interface{}) bool {
		node := v.(*CacheNode)
		if node.overTTL() {
			node.Destroy() // Destroy的回调会完成删除工作
		}
		return true
	})
}

// Close 关闭Cache，退出可能存在的TTL检查服务
func (cache *Cache) Close() {
	cache.Lock()
	defer cache.Unlock()
	if cache.destroyed != nil {
		close(cache.destroyed)
		cache.destroyed = nil
	}
}
