/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-27 20:46:32
 * @LastEditTime: 2019-03-27 22:55:25
 */

package trie

// RuneTrie Rune字典树
type RuneTrie struct {
	Trie
}

// Grow 生长
func (t *RuneTrie) Grow(data []rune) {
	t.Trie.Grow(runeSlice2ISlice(data))
}

// ReverseGrow 逆序生长
func (t *RuneTrie) ReverseGrow(data []rune) {
	t.Trie.ReverseGrow(runeSlice2ISlice(data))
}

// Find 查找
func (t RuneTrie) Find(data []rune) bool {
	return t.Trie.Find(runeSlice2ISlice(data))
}

// ReverseFind 逆序查找
func (t RuneTrie) ReverseFind(data []rune) bool {
	return t.Trie.ReverseFind(runeSlice2ISlice(data))
}

// MatchPrefix 匹配前缀
func (t RuneTrie) MatchPrefix(data []rune) bool {
	return t.Trie.MatchPrefix(runeSlice2ISlice(data))
}

// MatchSuffix 匹配后缀
func (t RuneTrie) MatchSuffix(data []rune) bool {
	return t.Trie.MatchSuffix(runeSlice2ISlice(data))
}

// Count 元素数量
func (t RuneTrie) Count() int {
	return t.Trie.Count()
}

// runeSlice2ISlice []rune -> []interface{}
func runeSlice2ISlice(data []rune) []interface{} {
	idata := make([]interface{}, len(data))
	for i := 0; i < len(data); i++ {
		idata[i] = data[i]
	}
	return idata
}
