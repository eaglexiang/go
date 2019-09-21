/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-02 17:28:48
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-09-21 14:09:04
 */

package tunnel

import (
	"net"
	"sync"
	"time"

	"github.com/eaglexiang/go/bytebuffer"
)

// Tunnel 是一个全双工的双向隧道，内置加密解密、暂停等待的控制器。
// 只能使用GetTunnel方法获取
type Tunnel struct {
	l          sync.Mutex
	left2Right *pipe
	right2Left *pipe
	flowed     bool
}

func newTunnel() *Tunnel {
	var t = Tunnel{
		left2Right: newPipe(),
		right2Left: newPipe(),
	}
	return &t
}

// Clear 清空Tunnel
func (t *Tunnel) Clear() {
	t.left2Right.Clear()
	t.right2Left.Clear()
}

// Left 左边的连接
func (t *Tunnel) Left() net.Conn {
	return t.left2Right.GetIn()
}

// Right 右边的连接
func (t *Tunnel) Right() net.Conn {
	return t.right2Left.GetIn()
}

// WriteLeft 向左边写数据
func (t *Tunnel) WriteLeft(b *bytebuffer.ByteBuffer) error {
	return t.right2Left.WriteOut(b)
}

// WriteRight 向右边写数据
func (t *Tunnel) WriteRight(b *bytebuffer.ByteBuffer) error {
	return t.left2Right.WriteOut(b)
}

// ReadLeft 从左边读取数据
func (t *Tunnel) ReadLeft(b *bytebuffer.ByteBuffer) error {
	return t.left2Right.ReadIn(b)
}

// ReadRight 从右边读取数据
func (t *Tunnel) ReadRight(b *bytebuffer.ByteBuffer) error {
	return t.right2Left.ReadIn(b)
}

// Close 关闭Tunnel，关闭前会停止其双向的流动
func (t *Tunnel) Close() (err error) {
	errl2r := t.left2Right.Close()
	errr2l := t.right2Left.Close()

	err = errl2r
	if errr2l != nil {
		err = errr2l
	}
	return
}

// Closed Tunnel是否已经关闭
func (t *Tunnel) Closed() bool {
	var closed = t.left2Right.Closed() && t.right2Left.Closed()
	return closed
}

// Flow 开始双向流动
// 此方法阻塞
// 同一个Tunnel同时只能执行一次Flow，多次Flow会导致panic
func (t *Tunnel) Flow() {
	t.l.Lock()
	if t.flowed {
		panic("Tunnel flowed already")
	}
	t.flowed = true
	t.l.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		t.right2Left.Flow()
		wg.Done()
	}()
	go func() {
		t.left2Right.Flow()
		wg.Done()
	}()
	wg.Wait()

	t.l.Lock()
	t.flowed = false
	t.l.Unlock()
}

// IsNil Tunnel的Left和Right都为nil
func (t *Tunnel) IsNil() bool {
	var left = t.Left()
	var right = t.Right()
	var rs = left == nil && right == nil
	return rs
}

// ReadLeftStr 从左边读字符串
func (t *Tunnel) ReadLeftStr() (text string, err error) {
	b := bytebuffer.GetBuffer()
	defer bytebuffer.PutBuffer(b)

	err = t.ReadLeft(b)
	text = b.String()
	return
}

// ReadRightStr 从左边读字符串
func (t *Tunnel) ReadRightStr() (text string, err error) {
	b := bytebuffer.GetBuffer()
	defer bytebuffer.PutBuffer(b)

	err = t.ReadRight(b)
	text = b.String()
	return
}

// WriteLeftStr 向左边写字符串
func (t *Tunnel) WriteLeftStr(str string) (err error) {
	b := bytebuffer.GetStringBuffer(str)
	defer bytebuffer.PutBuffer(b)

	err = t.WriteLeft(b)
	return
}

// WriteRightStr 向右边写字符串
func (t *Tunnel) WriteRightStr(str string) (err error) {
	b := bytebuffer.GetStringBuffer(str)
	defer bytebuffer.PutBuffer(b)

	err = t.WriteRight(b)
	return
}

// SetLeft 设置左边
func (t *Tunnel) SetLeft(conn net.Conn) {
	t.left2Right.In = conn
	t.right2Left.Out = conn
}

// SetRight 设置右边
func (t *Tunnel) SetRight(conn net.Conn) {
	t.left2Right.Out = conn
	t.right2Left.In = conn
}

// SetTimeout 设置超时时间
func (t *Tunnel) SetTimeout(timeout time.Duration) {
	t.left2Right.Timeout = timeout
	t.right2Left.Timeout = timeout
}
