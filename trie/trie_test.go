/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-27 19:23:56
 * @LastEditTime: 2019-03-27 20:55:43
 */

package trie

import "testing"

func Test_TrieFind(t *testing.T) {
	tr := Trie{}
	tr.Grow([]interface{}{1, 2, 3, 4, 5, 6})
	tr.Grow([]interface{}{4, 5, 6, 7, 8, 9})
	tr.Grow([]interface{}{2, 6, 8, 11, 5})
	tr.Grow([]interface{}{1, 4, 5, 6})

	if tr.Find([]interface{}{1, 2, 3, 4}) {
		t.Error("'1,2,3,4' shouldn't be found")
	}
	if tr.Find([]interface{}{}) {
		t.Error("nil []interface{} shouldn't be found")
	}
	if tr.Find([]interface{}{1, 4, 5, 6, 7}) {
		t.Error("'1,4,5,6,7' shouldn't be found")
	}
	if !tr.Find([]interface{}{1, 4, 5, 6}) {
		t.Error("'1,4,5,6' should be found")
	}
}

func Test_TrieMatch(t *testing.T) {
	tr := Trie{}
	tr.Grow([]interface{}{1, 2, 3, 4, 5, 6})
	tr.Grow([]interface{}{4, 5, 6, 7, 8, 9})
	tr.Grow([]interface{}{2, 6, 8, 11, 5})
	tr.Grow([]interface{}{1, 4, 5, 6})

	if tr.MatchPrefix([]interface{}{1, 2, 3, 4}) {
		t.Error("'1,2,3,4' shouldn't be matched")
	}
	if tr.MatchPrefix([]interface{}{}) {
		t.Error("nil shouldn't be matched")
	}
	if !tr.MatchPrefix([]interface{}{1, 4, 5, 6}) {
		t.Error("'1,4,5,6' should be matched")
	}
	if !tr.MatchPrefix([]interface{}{1, 4, 5, 6, 7}) {
		t.Error("'1,4,5,6,7' should be matched")
	}
}

func Test_Reverse(t *testing.T) {
	tr := Trie{}
	tr.ReverseGrow([]interface{}{4, 3, 2, 1})
	if !tr.Find([]interface{}{1, 2, 3, 4}) {
		t.Error("'1,2,3,4' should be found")
	}
	if !tr.MatchPrefix([]interface{}{1, 2, 3, 4}) {
		t.Error("'1,2,3,4' should be matched")
	}
	if !tr.MatchSuffix([]interface{}{4, 3, 2, 1}) {
		t.Error("'4,3,2,1' should be matched")
	}
}

func Test_reverse(t *testing.T) {
	data := []interface{}{1, 2, 3, 4}
	reverse(data)
	if data[0] != 4 {
		t.Error("failed to reverse")
	}
	if data[1] != 3 {
		t.Error("failed to reverse")
	}
	if data[2] != 2 {
		t.Error("failed to reverse")
	}
	if data[3] != 1 {
		t.Error("failed to reverse")
	}
}
