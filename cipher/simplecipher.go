/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-22 10:20:25
 * @LastEditTime: 2019-08-26 21:22:40
 */

package cipher

import (
	"errors"
	"strconv"

	"github.com/eaglexiang/go/bytebuffer"
)

// SimpleCipher 简单加密
type SimpleCipher struct {
	key byte
}

// SetKey 设置密码
func (sc *SimpleCipher) SetKey(key string) error {
	if key == "" {
		return errors.New("key is empty")
	}
	_key, err := strconv.ParseInt(key, 10, 8)
	if err != nil {
		return err
	}
	sc.key = byte(_key)
	return nil
}

// StreamEncrypt 加密
func (sc SimpleCipher) StreamEncrypt(input *bytebuffer.ByteBuffer) (output *bytebuffer.ByteBuffer) {
	output = input.Clone()
	data := output.Data()

	for i, value := range data {
		data[i] = value ^ sc.key
	}

	return
}

// StreamDecrypt 解密
func (sc SimpleCipher) StreamDecrypt(input *bytebuffer.ByteBuffer) (output *bytebuffer.ByteBuffer) {
	output = input.Clone()
	data := output.Data()

	for i, value := range data {
		data[i] = value ^ sc.key
	}

	return
}
