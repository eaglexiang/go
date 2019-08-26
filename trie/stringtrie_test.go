/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-27 22:34:21
 * @LastEditTime: 2019-03-27 22:58:55
 */

package trie

import "testing"

func Test_StringTriePrefix(t *testing.T) {
	st := StringTrie{}

	st.Grow("123456")
	st.Grow("34567")
	st.Grow("jdasfhi")
	st.Grow("123")

	if !st.MatchPrefix("1234") {
		t.Error("'1234' should be matched")
	}
	if !st.MatchPrefix("34567") {
		t.Error("'34567' should be matched")
	}
	if st.MatchPrefix("12") {
		t.Error("'12' shouldn't be matched")
	}
	if st.MatchPrefix(("")) {
		t.Error("'' shouldn't be matched")
	}
}

func Test_StringTrieSuffix(t *testing.T) {
	st := StringTrie{}

	st.ReverseGrow("exe")
	st.ReverseGrow("txt")
	st.ReverseGrow("jpg")
	st.ReverseGrow("mkv")

	if st.MatchSuffix("test.mp4") {
		t.Error("'mp4' shouldn't be matched")
	}
	if st.MatchSuffix("demo.mp3") {
		t.Error("'mp3' shouldn't be matched")
	}
	if !st.MatchSuffix("pic.jpg") {
		t.Error("'jpg' should be matched")
	}
	if !st.MatchSuffix("run.exe") {
		t.Error("'exe' should be matched")
	}
}

func Test_StringTrieCount(t *testing.T) {
	st := StringTrie{}

	st.Grow("12345")
	st.Grow("123")

	if st.Count() != 2 {
		t.Error("wrong count: ", st.Count())
	}

	st.Grow("123")

	if st.Count() != 2 {
		t.Error("wrong count: ", st.Count())
	}
}
