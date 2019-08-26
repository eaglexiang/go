/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-16 16:36:04
 * @LastEditTime: 2019-03-16 16:39:38
 */

package smartstrmap

import "testing"

func Test_SmartStrMap(t *testing.T) {
	m := SmartStrMap{}
	m.Bind("bindKey", "trueKey")
	m.Set("bindKey", "TrueValue")
	v, ok := m.Get("bindkey")
	if !ok {
		t.Error("bindkey not found")
	}
	if v != "TrueValue" {
		t.Error("value for bindkey is: ", v)
	}
	v, ok = m.Get("BINDKEY")
	if !ok {
		t.Error("BINDKEY not found")
	}
	if v != "TrueValue" {
		t.Error("value for BINDKEY is: ", v)
	}
	v, ok = m.Get("truekey")
	if !ok {
		t.Error("truekey not found")
	}
	if v != "TrueValue" {
		t.Error("value for truekey is: ", v)
	}
}
