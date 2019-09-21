/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-07-24 21:22:45
 * @LastEditTime: 2019-09-21 14:07:00
 */

package tunnel

import (
	"net"
	"sync"
	"time"

	"github.com/eaglexiang/go/bytebuffer"
)

// pipe 单向流动的数据管道
type pipe struct {
	l          *sync.Mutex
	In         net.Conn // 入口
	Out        net.Conn // 出口
	BufferSize int      // 数据缓冲区的尺寸
	Timeout    time.Duration
	flowed     bool
	closed     bool
}

func newPipe() *pipe {
	var p = new(pipe)
	p.l = new(sync.Mutex)
	return p
}

// Close 关闭Tunnel，关闭前会停止其流动
func (p *pipe) Close() (err error) {
	p.l.Lock()
	defer p.l.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	err = p.In.Close()
	return
}

// Closed 是否已经关闭
func (p *pipe) Closed() bool {
	return p.closed
}

// Clear 清空以便重新利用
func (p *pipe) Clear() {
	p.In = nil
	p.Out = nil
	p.closed = false
	p.flowed = false
	p.Timeout = 0
}

// In 获取管道的入口
func (p *pipe) GetIn() net.Conn {
	return p.In
}

// Out 获取管道的出口
func (p *pipe) GetOut() net.Conn {
	return p.Out
}

// Out 向出口写数据
func (p *pipe) WriteOut(b *bytebuffer.ByteBuffer) (err error) {
	return writePipeOut(b, p.Out, p.Timeout)
}

func writePipeOut(b *bytebuffer.ByteBuffer, conn net.Conn, timeout time.Duration) (err error) {
	if timeout != 0 {
		var dl = time.Now().Add(timeout)
		conn.SetWriteDeadline(dl)
	}

	b = b.Clone()
	defer bytebuffer.PutBuffer(b)
	for {
		n, err := conn.Write(b.Data())
		if err != nil {
			break
		}

		b.Move(n)
		if b.Length == 0 {
			break
		}
	}

	return
}

// ReadIn 从入口读数据
func (p *pipe) ReadIn(b *bytebuffer.ByteBuffer) (err error) {
	return readPipeIn(p.In, b, p.Timeout)
}

func readPipeIn(conn net.Conn, b *bytebuffer.ByteBuffer, timeout time.Duration) (err error) {
	if timeout != 0 {
		var dl = time.Now().Add(timeout)
		conn.SetDeadline(dl)
	}

	b.Length, err = conn.Read(b.Buf())
	return
}

// Flow 开始流动
// 数据会从入口流动到出口，并进行自动和透明的加解密处理
// 此方法阻塞
// 同一个pipe同时只能运行一次flow，多次flow会导致panic
func (p *pipe) Flow() {
	p.l.Lock()
	if p.flowed {
		panic("pipe flowed already")
	}
	p.flowed = true
	p.l.Unlock()

	p.flow()

	p.l.Lock()
	p.flowed = false
	p.l.Unlock()
}

// flow 开始流动
// 此方法阻塞
func (p *pipe) flow() {
	var b = make(chan *bytebuffer.ByteBuffer, p.BufferSize)
	go p.flowFromIn(b)
	p.flowToOut(b)
}

// flowFromIn 数据从入口流入
func (p *pipe) flowFromIn(bf chan *bytebuffer.ByteBuffer) {
	for {
		b := bytebuffer.GetBuffer()
		err := readPipeIn(p.In, b, p.Timeout)
		if err != nil {
			bytebuffer.PutBuffer(b)
			break
		}
		bf <- b
	}

	close(bf)
}

// flowToOut 数据从bf流入出口
func (p *pipe) flowToOut(bf chan *bytebuffer.ByteBuffer) {
	for b := range bf {
		err := writePipeOut(b, p.Out, p.Timeout)
		bytebuffer.PutBuffer(b)
		if err != nil {
			break
		}
	}
}
