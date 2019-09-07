/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-02 17:46:46
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-09-07 21:37:37
 */

package bytebuffer

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

var (
	pools       = make(map[int]*sync.Pool)
	defaultSize = 1000  // defaultSize 默认的Buffer尺寸
	debug       = false // debug开关，当debug为true，Put操作不会清空ByteBuffer
)

func init() {
	RegisterPool(defaultSize)
}

// DefaultSize 默认的Buffer尺寸
func DefaultSize() int {
	return defaultSize
}

// SetDefaultSize 设置默认的Buffer尺寸
func SetDefaultSize(size int) {
	defaultSize = size
	RegisterPool(defaultSize)
}

// RegisterPool 注册buffer池
func RegisterPool(size int) {
	if _, ok := pools[size]; !ok {
		pool := &sync.Pool{}
		pool.New = func() interface{} {
			return CreateByteBuffer(size)
		}
		pools[size] = pool
	}
}

func getBufferSize(size ...int) (_size int) {
	if len(size) == 0 {
		_size = defaultSize
	} else {
		_size = size[0]
	}

	return
}

// GetBuffer 尝试从Buffer池获取Buffer，如果size参数为空，尺寸为defaultSize
func GetBuffer(size ...int) (buffer *ByteBuffer) {
	var _size = getBufferSize(size...)

	if pool, ok := pools[_size]; ok {
		buffer = pool.Get().(*ByteBuffer)
	} else {
		buffer = CreateByteBuffer(_size)
	}

	return
}

// GetBytesBuffer 尝试从Buffer池获取Buffer，并拷入b的数据
func GetBytesBuffer(b []byte) (buffer *ByteBuffer) {
	buffer = GetBuffer(len(b))
	buffer.Length = copy(buffer.Buf(), b)
	return
}

// GetStringBuffer 尝试从Buffer池获取Buffer，并拷入str的数据
func GetStringBuffer(str string) (buffer *ByteBuffer) {
	b := []byte(str)
	buffer = GetBytesBuffer(b)
	return
}

// PutBuffer 归还Buffer，不符合默认尺寸的buffer会被丢弃
func PutBuffer(buffer *ByteBuffer) {
	l := buffer.Cap()
	if pool, ok := pools[l]; ok {
		if !debug {
			buffer.Clear()
		}
		pool.Put(buffer)
	}
}

// ByteBuffer bytes缓冲器
type ByteBuffer struct {
	buf    []byte
	Off    int
	Length int
}

// CreateByteBuffer 创建ByteBuffer
func CreateByteBuffer(cap int) *ByteBuffer {
	buffer := ByteBuffer{buf: make([]byte, cap)}
	return &buffer
}

// Buf 获取buf
func (buffer ByteBuffer) Buf() []byte {
	return buffer.buf[buffer.Off:]
}

// Cap 获取容量
func (buffer ByteBuffer) Cap() int {
	return len(buffer.Buf())
}

// Clear 清空
func (buffer *ByteBuffer) Clear() {
	buffer.Off = 0
	buffer.Length = 0
}

// String 字符串
func (buffer ByteBuffer) String() string {
	var b = buffer.Data()
	var str = string(b)

	return str
}

// SetString 将Buffer的值设置为字符串
func (buffer *ByteBuffer) SetString(data string) (err error) {
	var b = []byte(data)
	if len(b) > buffer.Cap() {
		err = errors.New("buffer overflow")
		return
	}

	buffer.Length = copy(buffer.Buf(), b)
	return
}

// Cut 返回实际数据部分的[]byte，会构造新的数组
func (buffer ByteBuffer) Cut() []byte {
	start := buffer.Off
	end := buffer.Off + buffer.Length
	oldBuf := buffer.buf[start:end]

	newBuf := make([]byte, buffer.Length)

	copy(newBuf, oldBuf)

	return newBuf
}

// Data 返回实际数据部分的[]byte，不会构造新的数组
func (buffer ByteBuffer) Data() []byte {
	start := buffer.Off
	end := buffer.Off + buffer.Length

	return buffer.buf[start:end]
}

// Clone 克隆一个新的ByteBuffer
func (buffer ByteBuffer) Clone() (newBuffer *ByteBuffer) {
	newBuffer = GetBuffer(buffer.Length)
	newBuffer.Off = buffer.Off
	newBuffer.Length = copy(newBuffer.Buf(), buffer.Data())
	return
}

// Move 将数据部分的起点后移step距离，当step为负数，则为前移
func (buffer *ByteBuffer) Move(step int) (err error) {
	off := buffer.Off + step
	length := buffer.Length - step

	if off > buffer.Cap() || length > buffer.Cap() || off < 0 || length < 0 {
		errStr := fmt.Sprint("invalid buffer after move. cap: ", buffer.Cap(), ", off: ", buffer.Off, ", length: ", buffer.Length)
		err = errors.New(errStr)
		return
	}

	buffer.Off = off
	buffer.Length = length
	return
}

// MarshalJSON JSON序列化器
func (buffer *ByteBuffer) MarshalJSON() (b []byte, err error) {
	data := make(map[string]interface{})
	data["off"] = buffer.Off
	data["length"] = buffer.Length
	data["buf"] = buffer.buf

	b, err = json.Marshal(data)
	return
}

// UnmarshalJSON JSON反序列化器
func (buffer *ByteBuffer) UnmarshalJSON(b []byte) (err error) {
	data := make(map[string]interface{})
	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}

	if _off, ok := data["off"]; ok {
		if off, ok := _off.(float64); ok {
			buffer.Off = int(off)
		} else {
			err = errors.New("type for Off should be float64 but " + fmt.Sprintf("%T", _off))
			return
		}
	}

	if _length, ok := data["length"]; ok {
		if length, ok := _length.(float64); ok {
			buffer.Length = int(length)
		} else {
			err = errors.New("type for Length should be float64 but " + fmt.Sprintf("%T", _length))
			return
		}
	}

	if _buf, ok := data["buf"]; ok {
		if buf, ok := _buf.(string); ok {
			buffer.buf = []byte(buf)
		} else {
			err = errors.New("type for buf should be string but " + fmt.Sprintf("%T", _buf))
			return
		}
	}

	return
}
