/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-27 17:55:10
 * @LastEditTime: 2019-03-27 22:52:01
 */

package trie

type trieNode struct {
	exsitTail bool
	children  map[interface{}]*trieNode
}

func (n *trieNode) TryAdd(i interface{}, isTail bool) (child *trieNode, newTail bool) {
	if n.children == nil {
		n.children = make(map[interface{}]*trieNode)
	}

	var ok bool
	child, ok = n.children[i]
	if !ok {
		child = new(trieNode)
		n.children[i] = child
	}
	if isTail {
		if !child.exsitTail {
			newTail = true
		}
		child.exsitTail = true
	}
	return
}

func (n trieNode) Find(i interface{}) (child *trieNode) {
	if n.children == nil {
		return nil
	}

	return n.children[i]
}
