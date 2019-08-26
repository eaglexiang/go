/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-04-01 20:35:23
 * @LastEditTime: 2019-04-01 21:27:30
 */

package net

import "testing"

func Test_CheckIPType(t *testing.T) {
	if TypeOfAddr("192.168.50.10") != IPv4Addr {
		t.Error("192.168.50.10 should be IPv4")
	}
	if TypeOfAddr("::") != IPv6Addr {
		t.Error(":: should be IPv6")
	}
	if TypeOfAddr("test.com") != DomainAddr {
		t.Error("test.com may be domain")
	}
}

func Test_CheckPrivateIPv4(t *testing.T) {
	if !IsPrivateIPv4("127.0.0.1") {
		t.Error("127.0.0.1 is local loop")
	}
	if !IsPrivateIPv4("192.168.0.1") {
		t.Error("192.168.0.1 is private addr")
	}
	if !IsPrivateIPv4("172.16.0.1") {
		t.Error("172.16.0.1 is private addr")
	}
	if IsPrivateIPv4("171.217.167.174") {
		t.Error("171.217.167.174 is not private addr")
	}
}
