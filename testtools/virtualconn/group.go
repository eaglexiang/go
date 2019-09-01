/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-09-01 11:28:10
 * @LastEditTime: 2019-09-01 11:32:15
 */

package virtualconn

// NewGroup 构建虚拟连接小组，左边和右边相互联通
func NewGroup() (left, right *VirtualConn) {
	left = New()
	right = New()

	go func() {
		for {
			b := left.GetWriteBuf()
			right.PutReadBuf(b)
		}
	}()

	go func() {
		for {
			b := right.GetWriteBuf()
			left.PutReadBuf(b)
		}
	}()

	return
}
