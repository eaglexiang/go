<!--
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-21 16:37:16
 * @LastEditTime: 2019-06-14 20:35:28
 -->
 
# go-bytebuffer

Byte Buffer struct for Go

## 自带对象池

使用实例

```golang
func GetBuffer() *ByteBuffer // 获取容量为Buffer
func PutBuffer(*ByteBuffer)  // 归还容量为Buffer

func SetDefaultSize(int)     // 设置Buffer的尺寸
func DefaultSize() int       // 获取当前Buffer的尺寸
```

归还错误尺寸的Buffer不会造成异常，错误的Buffer会被丢弃
