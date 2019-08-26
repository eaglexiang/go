/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-07-24 21:22:31
 * @LastEditTime: 2019-07-24 21:22:33
 */

package tunnel

import "sync"

// TunnelPool *Tunnel Pool
var tunnelPool sync.Pool

func init() {
	tunnelPool.New = func() interface{} {
		return newTunnel()
	}
}

// GetTunnel 从Tunnel Pool获取*Tunnel
func GetTunnel() *Tunnel {
	tunnel := tunnelPool.Get().(*Tunnel)
	return tunnel
}

// PutTunnel 将*Tunnel放回TunnelPool
func PutTunnel(tunnel *Tunnel) {
	tunnel.Clear()
	tunnelPool.Put(tunnel)
}
