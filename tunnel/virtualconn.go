/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-06-02 10:49:54
 * @LastEditTime: 2019-12-07 17:17:58
 */

package tunnel

import (
	"io"
	"net"
	"time"
)

// NewVirtualConn 创建新的虚拟连接
func NewVirtualConn() net.Conn {
	return &virtualConn{
		msgs: make(chan []byte, 2),
	}
}

type virtualConn struct {
	msgs chan []byte
}

func (conn virtualConn) Write(b []byte) (n int, err error) {
	conn.msgs <- b
	return len(b), nil
}

func (conn virtualConn) Read(b []byte) (n int, err error) {
	tmp, ok := <-conn.msgs
	if !ok {
		return 0, io.EOF
	}

	copy(b, tmp)
	return len(tmp), nil
}

func (conn virtualConn) Close() (err error) {
	close(conn.msgs)
	return nil
}

func (conn virtualConn) LocalAddr() (addr net.Addr) { return nil }

func (conn virtualConn) RemoteAddr() (addr net.Addr) { return nil }

func (conn virtualConn) SetWriteDeadline(t time.Time) error { return nil }

func (conn virtualConn) SetReadDeadline(t time.Time) error { return nil }

func (conn virtualConn) SetDeadline(t time.Time) error { return nil }
