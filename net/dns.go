/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-22 11:11:08
 * @LastEditTime: 2019-04-01 20:35:05
 */

package net

import (
	"errors"
	"net"
)

// AddrType 地址的类型
type AddrType int

// 地址类型
const (
	InvalidAddr AddrType = iota
	DomainAddr
	IPv4Addr
	IPv6Addr
)

// ResolvIPv4 本地解析IPv4
func ResolvIPv4(domain string) (ip string, err error) {
	addrs, err := net.LookupHost(domain)
	if err != nil {
		return "", errors.New("ResolvIPv4 -> " + err.Error())
	}
	for _, addr := range addrs {
		if TypeOfAddr(addr) == IPv4Addr {
			return addr, nil
		}
	}
	return "", errors.New("ResolvIPv4 -> not found")
}

// ResolvIPv6 本地解析IPv6
func ResolvIPv6(domain string) (ip string, err error) {
	addrs, err := net.LookupHost(domain)
	if err != nil {
		return "", errors.New("ResolvIPv6 -> " + err.Error())
	}
	for _, addr := range addrs {
		if TypeOfAddr(addr) == IPv6Addr {
			return addr, nil
		}
	}
	return "", errors.New("ResolvIPv6 -> not found")
}

// TypeOfAddr 地址的类型
func TypeOfAddr(host string) AddrType {
	ip := net.ParseIP(host)
	if ip == nil {
		return DomainAddr
	}
	switch {
	case ip.To4() == nil && ip.To16() != nil:
		return IPv6Addr
	case ip.To4() != nil:
		return IPv4Addr
	default:
		return InvalidAddr
	}
}
