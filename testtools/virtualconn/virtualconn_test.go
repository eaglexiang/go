/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-17 10:41:50
 * @LastEditTime: 2019-08-17 10:52:34
 */

package virtualconn

import (
	"testing"
)

func Test_Write(t *testing.T) {
	conn := New()

	conn.Write([]byte("test msg"))
	buf := conn.GetWriteBuf()
	if string(buf) != "test msg" {
		t.Error("wrong wirte buf: ", string(buf))
	}
}

func Test_Read(t *testing.T) {
	conn := New()

	conn.PutReadBuf([]byte("test msg"))
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	if string(buf[:n]) != "test msg" {
		t.Error("wrong buf: ", string(buf[:n]))
	}
}
