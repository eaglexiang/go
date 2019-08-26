/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-20 21:12:15
 * @LastEditTime: 2019-08-26 21:22:55
 */

package cipher

import (
	"github.com/eaglexiang/go/bytebuffer"
)

// StreamEncryptor 流式加密器
type StreamEncryptor interface {
	StreamEncrypt(*bytebuffer.ByteBuffer) *bytebuffer.ByteBuffer
}

// StreamDecryptor 流式解密器
type StreamDecryptor interface {
	StreamDecrypt(*bytebuffer.ByteBuffer) *bytebuffer.ByteBuffer
}
