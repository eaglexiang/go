/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-17 10:10:15
 * @LastEditTime: 2019-08-26 23:35:24
 */

package cipherconn

import (
	"net"

	"github.com/eaglexiang/go/bytebuffer"
	mycipher "github.com/eaglexiang/go/cipher"
)

// connect 加密连接
type connect struct {
	net.Conn                          // 基础连接
	cIn      mycipher.StreamDecryptor // Input Cipher
	cOut     mycipher.StreamEncryptor // Output Cipher
}

func (conn connect) Read(b []byte) (n int, err error) {
	if conn.cIn != nil {
		n, err = conn.readEncrypted(b)
	} else {
		n, err = conn.read(b)
	}
	return
}

func (conn connect) readEncrypted(b []byte) (n int, err error) {
	buf := bytebuffer.GetBuffer(len(b))
	defer bytebuffer.PutBuffer(buf)

	buf.Length, err = conn.Conn.Read(buf.Buf())

	newBuf := conn.cIn.StreamDecrypt(buf)
	defer bytebuffer.PutBuffer(newBuf)

	n = copy(b, newBuf.Data())
	return
}

func (conn connect) read(b []byte) (n int, err error) {
	n, err = conn.Conn.Read(b)
	return
}

func (conn connect) Write(b []byte) (n int, err error) {
	if conn.cOut != nil {
		n, err = conn.writeEncryped(b)
	} else {
		n, err = conn.write(b)
	}
	return
}

func (conn connect) writeEncryped(b []byte) (n int, err error) {
	buf := bytebuffer.GetBytesBuffer(b)
	defer bytebuffer.PutBuffer(buf)

	newBuf := conn.cOut.StreamEncrypt(buf)
	defer bytebuffer.PutBuffer(newBuf)

	n, err = conn.Conn.Write(newBuf.Cut())
	return
}

func (conn connect) write(b []byte) (n int, err error) {
	n, err = conn.Conn.Write(b)
	return
}

// New 构造新的CipherConnect
func New(base net.Conn, cIn mycipher.StreamDecryptor, cOut mycipher.StreamEncryptor) net.Conn {
	var conn = &connect{
		Conn: base,
		cIn:  cIn,
		cOut: cOut,
	}
	return conn
}
