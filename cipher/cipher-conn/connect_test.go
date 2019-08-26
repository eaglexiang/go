/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-17 10:30:56
 * @LastEditTime: 2019-08-26 23:35:35
 */

package cipherconn

import (
	"testing"

	"github.com/eaglexiang/go/bytebuffer"
	"github.com/eaglexiang/go/cipher"
	"github.com/eaglexiang/go/testtools/virtualconn"
)

func Test_Read(t *testing.T) {
	base := virtualconn.New()
	c := cipher.SimpleCipher{}
	c.SetKey("55")
	conn := New(base, c, nil)

	in := "test in"
	b := bytebuffer.GetBytesBuffer([]byte(in))
	b = c.StreamEncrypt(b)
	base.PutReadBuf(b.Data())

	_in := make([]byte, 1024)
	n, _ := conn.Read(_in)

	_In := string(_in[:n])
	if _In != "test in" {
		t.Error("_In should be test in but: ", _In)
	}
}

func Test_Write(t *testing.T) {
	base := virtualconn.New()
	c := cipher.SimpleCipher{}
	c.SetKey("55")
	conn := New(base, nil, c)

	out := "test out"
	conn.Write([]byte(out))

	_out := base.GetWriteBuf()
	b := bytebuffer.GetBytesBuffer(_out)
	b = c.StreamDecrypt(b)

	_Out := string(b.Data())
	if _Out != "test out" {
		t.Error("_Out should be test out but: ", _Out)
	}
}
