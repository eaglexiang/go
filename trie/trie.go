/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-27 19:19:43
 * @LastEditTime: 2019-03-27 22:53:29
 */

package trie

// Trie 字典树
type Trie struct {
	root  *trieNode
	count int // 总元素数量
}

// Grow 根据给定的[]interface{}进行生长
func (t *Trie) Grow(data []interface{}) {
	if t.root == nil {
		t.root = new(trieNode)
	}

	var (
		nodeNow = t.root
		newTail bool
	)
	for i := 0; i < len(data); i++ {
		nodeNow, newTail = nodeNow.TryAdd(data[i], i+1 == len(data))
	}
	if newTail {
		t.count++
	}
}

// ReverseGrow 根据给定[]interface{}的倒序进行生长
func (t *Trie) ReverseGrow(data []interface{}) {
	reverse(data)
	t.Grow(data)
}

// Find 在Trie中查找该数组是否存在
func (t Trie) Find(data []interface{}) bool {
	if t.root == nil {
		return false
	}

	nodeNow := t.root
	for i := 0; i < len(data); i++ {
		if nodeNow = nodeNow.Find(data[i]); nodeNow == nil {
			return false
		}
	}
	return nodeNow.exsitTail
}

// ReverseFind 在Trie中倒序查找数组
func (t Trie) ReverseFind(data []interface{}) bool {
	reverse(data)
	return t.Find(data)
}

// MatchPrefix 在Trie中查找是否存在该数组匹配的前缀
func (t Trie) MatchPrefix(data []interface{}) bool {
	if t.root == nil {
		return true
	}

	nodeNow := t.root
	for i := 0; i < len(data); i++ {
		nodeNow = nodeNow.Find(data[i])
		if nodeNow == nil {
			return false
		}
		if nodeNow.exsitTail {
			return true
		}
	}
	return false
}

// MatchSuffix 在Trie中查找是否存在该数组匹配的后缀
func (t Trie) MatchSuffix(data []interface{}) bool {
	reverse(data)
	return t.MatchPrefix(data)
}

// Count 数量
func (t Trie) Count() int {
	return t.count
}

func reverse(data []interface{}) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
}
