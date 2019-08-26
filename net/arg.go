/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-22 15:26:39
 * @LastEditTime: 2019-08-26 23:36:56
 */

package net

import tunnel "github.com/eaglexiang/go/tunnel"

// Arg 网络业务会用到的参数集
type Arg struct {
	Msg       []byte         // 消息
	Host      string         // 主机地址:端口
	Tunnel    *tunnel.Tunnel // 数据隧道
	TheType   int            // 业务类型
	Delegates []func() bool  // 委托队列
}

// OpType 网络操作的类型
type OpType int

// 网络操作类型
const (
	ERROR OpType = iota
	CONNECT
	BIND
	UDP
)

// ResultOfNetOp 网络操作的结果
type ResultOfNetOp int

// 网络操作是否成功
const (
	FAILED ResultOfNetOp = iota
	SUCCESS
)
