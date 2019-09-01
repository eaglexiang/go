/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-09-01 20:12:38
 * @LastEditTime: 2019-09-01 20:29:20
 */

package queue

import (
	"container/list"
	"errors"

	"github.com/eaglexiang/go/bytebuffer"
)

// Queue 供ByteBuffer使用的队列
type Queue struct {
	data list.List
	size int
}

// New 构造新队列
func New(size ...int) (q *Queue) {
	q = new(Queue)
	if len(size) > 0 {
		q.size = size[0]
	}

	return
}

// Enqueue 入队
func (q *Queue) Enqueue(b *bytebuffer.ByteBuffer) (err error) {
	if q.size > 0 && q.data.Len() == q.size {
		err = errors.New("queue is full")
		return
	}
	if q.size > 0 && q.data.Len() > q.size {
		panic("queue's len is greater than size")
	}

	q.data.PushBack(b)
	return
}

// Dequeue 出队
func (q *Queue) Dequeue() (b *bytebuffer.ByteBuffer) {
	front := q.data.Front()
	if front == nil {
		return nil
	}

	v := q.data.Remove(front)
	b = v.(*bytebuffer.ByteBuffer)
	return
}
