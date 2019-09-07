/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-20 20:49:26
 * @LastEditTime: 2019-09-07 21:37:56
 */
package bytebuffer

import (
	"encoding/json"
	"testing"
)

func Test_SetString(t *testing.T) {
	b := GetBuffer(100)
	b.SetString("test")
	if b.String() != "test" {
		t.Error("b.String() should be test but ", b.String())
	}
}

func Test_Register(t *testing.T) {
	debug = true

	b := GetBuffer(100)
	b.SetString("test")
	PutBuffer(b)
	b = GetBuffer(100)
	if b.String() == "test" {
		t.Error("100 size pool not created")
	}

	RegisterPool(100)
	b.SetString("test")
	PutBuffer(b)
	b = GetBuffer(100)
	if b.String() != "test" {
		t.Error("100 size should be created now")
	}
}

func Test_Clone(t *testing.T) {
	b := GetBuffer(1000)
	b.SetString("test")
	_b := b.Clone()
	b.SetString("test 1")

	if _b.String() != "test" {
		t.Error("_b.String() should be test but ", _b.String())
	}
}

func Test_Move(t *testing.T) {
	b := GetBuffer()
	b.SetString("test")

	b.Move(1)
	if b.String() != "est" {
		t.Error("wrong b string: ", b.String())
	}
	b.Move(-1)
	if b.String() != "test" {
		t.Error("wrong b string: ", b.String())
	}
}

func Test_MarshalJSON(t *testing.T) {
	b := GetStringBuffer("test")

	j, err := json.Marshal(b)
	if err != nil {
		t.Error(err)
	}

	_b := new(ByteBuffer)
	err = json.Unmarshal(j, _b)
	if err != nil {
		t.Error(err)
	}

	if b.String() != "test" {
		t.Error("wrong b.String: ", b.String())
	}
}
