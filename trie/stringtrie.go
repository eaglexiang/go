/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-27 21:11:56
 * @LastEditTime: 2019-03-27 22:56:04
 */

package trie

// StringTrie String字典
type StringTrie struct {
	RuneTrie
}

// Grow 生长
func (t *StringTrie) Grow(s string) {
	t.RuneTrie.Grow([]rune(s))
}

// ReverseGrow 倒序生长
func (t *StringTrie) ReverseGrow(s string) {
	t.RuneTrie.ReverseGrow([]rune(s))
}

// Find 查找
func (t StringTrie) Find(s string) bool {
	return t.RuneTrie.Find([]rune(s))
}

// ReverseFind 倒序查找
func (t StringTrie) ReverseFind(s string) bool {
	return t.RuneTrie.ReverseFind([]rune(s))
}

// MatchPrefix 匹配前缀
func (t StringTrie) MatchPrefix(s string) bool {
	return t.RuneTrie.MatchPrefix([]rune(s))
}

// MatchSuffix 匹配后缀
func (t StringTrie) MatchSuffix(s string) bool {
	return t.RuneTrie.MatchSuffix([]rune(s))
}

// Count 元素数量
func (t StringTrie) Count() int {
	return t.RuneTrie.Count()
}
