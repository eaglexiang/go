/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-27 16:23:36
 * @LastEditTime: 2019-08-28 20:27:58
 */

package settings

import "testing"

func Test_Settings(t *testing.T) {
	Bind("testKey", "TrueKey")
	SetDefault("testKey", "testValueDefault")
	testValue := Get("testKey")
	if testValue != "testValueDefault" {
		t.Error("value must be testValueDefault")
	}
	Set("testKey", "testValue")
	testValue = Get("testKey")
	if testValue != "testValue" {
		t.Error("value must be testValue")
	}
	testValue = Get("testkey")
	if testValue != "testValue" {
		t.Error("value must be testValue")
	}
	testValue = Get("TESTkEY")
	if testValue != "testValue" {
		t.Error("value must be testValue")
	}
	testValue = Get("truekey")
	if testValue != "testValue" {
		t.Error("value must be testValue")
	}
	SetDefault("testKey", "testValueDefault")
	testValue = Get("testKey")
	if testValue == "testValueDefault" {
		t.Error("value mustn't be testValueDefault")
	}
}

func Test_GetChild(t *testing.T) {
	Set("parent.getChildTestKey", "getChildValue")
	child := GetChild("parent")
	value := child.Get("getChildTestKey")

	if value != "getChildValue" {
		t.Error("value for getChildTestKey should be gitChildValue, but ", value)
	}
}
