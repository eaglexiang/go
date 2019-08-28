/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-26 22:29:35
 * @LastEditTime: 2019-08-28 20:22:43
 */

package settings

import (
	"testing"
)

func Test_getChild(t *testing.T) {
	_, _, ok := getChild("testKey")
	if ok {
		t.Error("testKey should not has className")
	}

	className, subKey, ok := getChild("testParent.testKey")
	if !ok {
		t.Error("testParent.testKey should has className testParent")
	}
	if className != "testParent" {
		t.Error("testParent.testKey should has className testParent but ", className)
	}
	if subKey != "testKey" {
		t.Error("testParent.testKey should has subKey testKey but ", subKey)
	}
}
