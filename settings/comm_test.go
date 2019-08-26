/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-26 22:29:35
 * @LastEditTime: 2019-08-26 22:39:37
 */

package settings

import (
	"testing"
)

func Test_getClassName(t *testing.T) {
	_, ok := getClassName("testKey")
	if ok {
		t.Error("testKey should not has className")
	}

	className, ok := getClassName("testParent.testKey")
	if !ok {
		t.Error("testParent.testKey should has className testParent")
	}
	if className != "testParent" {
		t.Error("testParent.testKey should has className testParent but ", className)
	}
}

func Test_getSubKey(t *testing.T) {
	_, ok := getSubKey("testKey")
	if ok {
		t.Error("testKey should not has subKey")
	}

	subKey, ok := getSubKey("testParent.testKey")
	if !ok {
		t.Error("testParent.testKey should has subKey testKey")
	}
	if subKey != "testKey" {
		t.Error("testParent.testKey should has subKey testKey but ", subKey)
	}
}
