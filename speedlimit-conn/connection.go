/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-25 12:37:53
 * @LastEditTime: 2019-08-25 13:07:27
 */

package speedlimitconn

import (
	"math"
	"net"

	"go.uber.org/ratelimit"
)

const (
	limitPerBytes = 1000 // KiB
	limiterSize   = 1000
)

// speedLimitConn 带速度限制的Conn
type speedLimitConn struct {
	net.Conn
	l ratelimit.Limiter
}

// New 构造新的SpeedLimitConn
func New(base net.Conn, l ratelimit.Limiter) net.Conn {
	conn := speedLimitConn{
		Conn: base,
		l:    l,
	}

	return conn
}

func (conn speedLimitConn) Write(b []byte) (n int, err error) {
	n, err = conn.Conn.Write(b)
	if err != nil {
		return
	}

	if conn.l == nil {
		return
	}

	count := float64(n) / limitPerBytes
	count = math.Ceil(count)
	for i := float64(0); i < count; i++ {
		conn.l.Take()
	}

	return
}
