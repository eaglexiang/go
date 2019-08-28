/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-07-24 21:22:45
 * @LastEditTime: 2019-08-28 19:39:06
 */

package tunnel

import (
	"net"
	"time"

	"github.com/eaglexiang/go/bytebuffer"
	"github.com/pkg/errors"
)

// pipe 单向流动的数据管道，可透明进行数据的加解密，以及限速工作。
type pipe struct {
	in      net.Conn // 入口
	out     net.Conn // 出口
	bSize   int      // 数据缓冲区的尺寸
	timeout time.Duration
	pipeStatus
}

func newPipe() *pipe {
	var p = new(pipe)
	return p
}

// SetIn 设置入口
func (p *pipe) SetIn(in net.Conn) {
	p.in = in
}

// SetOut 设置出口
func (p *pipe) SetOut(out net.Conn) {
	p.out = out
}

// SetBufferSize 设置buffer的尺寸
func (p *pipe) SetBufferSize(size int) {
	p.bSize = size
}

// SetTimeout 设置超时时间
func (p *pipe) SetTimeout(timeout time.Duration) {
	p.timeout = timeout
}

// Close 关闭Tunnel，关闭前会停止其流动
func (p *pipe) Close() {
	p.l.Lock()
	defer p.l.Unlock()

	p.pipeStatus.Close()
	p.in.Close()
}

// Closed 是否已经关闭
func (p *pipe) Closed() bool {
	return p.pipeStatus.Closed()
}

// Clear 清空以便重新利用
func (p *pipe) Clear() {
	p.l.Lock()
	defer p.l.Unlock()

	p.in = nil
	p.out = nil
	p.pipeStatus.Clear()
	p.timeout = 0
}

// In 获取管道的入口
func (p *pipe) GetIn() net.Conn {
	return p.in
}

// Out 获取管道的出口
func (p *pipe) GetOut() net.Conn {
	return p.out
}

// Out 向出口写数据
func (p *pipe) WriteOut(b *bytebuffer.ByteBuffer) (err error) {
	p.l.Lock()
	var out = p.out
	var timeout = p.timeout
	p.l.Unlock()

	return writePipeOut(b, out, timeout)
}

func writePipeOut(b *bytebuffer.ByteBuffer, conn net.Conn, timeout time.Duration) (err error) {
	if conn == nil {
		err = errors.New("out of pipe is nil")
		return
	}

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
	p.l.Lock()
	var in = p.in
	var timeout = p.timeout
	p.l.Unlock()

	return readPipeIn(in, b, timeout)
}

func readPipeIn(conn net.Conn, b *bytebuffer.ByteBuffer, timeout time.Duration) (err error) {
	if conn == nil {
		err = errors.New("in of pipe is nil")
		return
	}

	if timeout != 0 {
		var dl = time.Now().Add(timeout)
		conn.SetDeadline(dl)
	}

	b.Length, err = conn.Read(b.Buf())
	if err != nil {
		return
	}

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

	defer func() {
		p.l.Lock()
		p.flowed = false
		p.l.Unlock()
	}()

	p.flow()
}

// flow 开始流动
// 此方法阻塞
func (p *pipe) flow() {
	var b = make(chan *bytebuffer.ByteBuffer, p.bSize)
	go p.flowFromIn(b)
	p.flowToOut(b)
}

// flowFromIn 数据从入口流入
func (p *pipe) flowFromIn(bf chan *bytebuffer.ByteBuffer) {
	p.l.Lock()
	var (
		in      = p.in
		timeout = p.timeout
	)
	p.l.Unlock()

	for {
		b := bytebuffer.GetBuffer()
		err := readPipeIn(in, b, timeout)
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
	p.l.Lock()
	var out = p.out
	var timeout = p.timeout
	p.l.Unlock()

	for b := range bf {
		err := writePipeOut(b, out, timeout)
		bytebuffer.PutBuffer(b)
		if err != nil {
			break
		}
	}
}
