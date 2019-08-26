/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-03 03:39:48
 * @LastEditTime: 2019-08-26 23:38:17
 */

package textcache

import (
	mycache "github.com/eaglexiang/go/cache"
)

// CacheNode 对cache.CacheNode的封装
// 用来等待已有的DNS请求解阻塞
type CacheNode struct {
	node *mycache.CacheNode
}

// Wait 等待DNS解析请求的返回
func (node CacheNode) Wait() (value string, err error) {
	if v, err := node.node.Wait4Value(); err == nil {
		value = v.(string)
	}
	return
}

// Update 更新值
// 如果该记录处于阻塞状态，更新操作会解除阻塞
func (node CacheNode) Update(value string) {
	node.node.Update(value)
}

// Destroy 销毁节点
func (node CacheNode) Destroy() {
	node.node.Destroy()
}
