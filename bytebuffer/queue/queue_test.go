/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-09-01 20:23:53
 * @LastEditTime: 2019-09-01 20:29:50
 */

package queue

import (
	"testing"

	"github.com/eaglexiang/go/bytebuffer"
)

func Test_Enqueue_Dequeue(t *testing.T) {
	a := bytebuffer.GetStringBuffer("testA")
	b := bytebuffer.GetStringBuffer("testB")

	q := New()
	q.Enqueue(a)
	q.Enqueue(b)
	a = q.Dequeue()
	b = q.Dequeue()
	c := q.Dequeue()

	if a.String() != "testA" {
		t.Error("a should be testA but ", a.String())
	}
	if b.String() != "testB" {
		t.Error("b should be testB but ", b.String())
	}
	if c != nil {
		t.Error("c should be nil")
	}
}
