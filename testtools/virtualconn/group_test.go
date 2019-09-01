/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-09-01 11:32:29
 * @LastEditTime: 2019-09-01 11:34:06
 */

package virtualconn

import "testing"

func Test_NewGroup(t *testing.T) {
	left, right := NewGroup()
	left.Write([]byte("test"))

	buf := make([]byte, 1024)
	n, _ := right.Read(buf)
	r := string(buf[:n])

	if r != "test" {
		t.Error("failed to transfer data test")
	}
}
