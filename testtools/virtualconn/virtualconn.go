/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-17 10:35:34
 * @LastEditTime: 2019-09-07 20:48:03
 */

package virtualconn

import (
	"net"
	"time"
)

// VirtualConn 虚拟连接
type VirtualConn struct {
	bufRead  chan []byte
	bufWrite chan []byte
}

func (conn VirtualConn) Read(b []byte) (n int, err error) {
	buf := <-conn.bufRead
	n = copy(b, buf)
	return
}

func (conn VirtualConn) Write(b []byte) (n int, err error) {
	buf := make([]byte, len(b))
	n = copy(buf, b)
	conn.bufWrite <- buf
	return
}

// PutReadBuf 将buf投入Read缓冲区
func (conn VirtualConn) PutReadBuf(b []byte) {
	buf := make([]byte, len(b))
	copy(buf, b)
	conn.bufRead <- buf
}

// GetWriteBuf 从Write缓冲区取出buf
func (conn VirtualConn) GetWriteBuf() []byte {
	b := <-conn.bufWrite
	buf := make([]byte, len(b))
	copy(buf, b)
	return buf
}

// Close 关闭
func (conn VirtualConn) Close() (err error) { return }

// LocalAddr 本地地址
func (conn VirtualConn) LocalAddr() (addr net.Addr) { return }

// RemoteAddr 远程地址
func (conn VirtualConn) RemoteAddr() (addr net.Addr) { return }

// SetDeadline 设置dl
func (conn VirtualConn) SetDeadline(t time.Time) (err error) { return }

// SetReadDeadline 设置读dl
func (conn VirtualConn) SetReadDeadline(t time.Time) (err error) { return }

// SetWriteDeadline 设置写dl
func (conn VirtualConn) SetWriteDeadline(t time.Time) (err error) { return }

// New 构造新的虚拟连接
func New() *VirtualConn {
	var conn = &VirtualConn{
		bufRead:  make(chan []byte, 1),
		bufWrite: make(chan []byte, 1),
	}
	return conn
}
