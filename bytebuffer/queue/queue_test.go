/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-09-01 20:23:53
 * @LastEditTime: 2019-09-01 20:44:39
 */

package queue

import (
	"sync"
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

// 测试并发情况下是否会触发panic
func Test_Sync(t *testing.T) {
	q := New(5)
	wg := sync.WaitGroup{}
	wg.Add(8)

	go func() {
		for i := 0; i < 10000; i++ {
			q.Enqueue(bytebuffer.GetStringBuffer("1"))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			q.Enqueue(bytebuffer.GetStringBuffer("1"))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			q.Enqueue(bytebuffer.GetStringBuffer("1"))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			q.Enqueue(bytebuffer.GetStringBuffer("1"))
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			q.Dequeue()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			q.Dequeue()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			q.Dequeue()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			q.Dequeue()
		}
		wg.Done()
	}()

	wg.Wait()
}
