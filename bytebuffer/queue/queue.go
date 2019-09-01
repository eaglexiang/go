/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-09-01 20:12:38
 * @LastEditTime: 2019-09-01 20:40:10
 */

package queue

import (
	"container/list"
	"sync"

	"github.com/eaglexiang/go/bytebuffer"
)

// Queue 供ByteBuffer使用的队列
type Queue struct {
	data   list.List
	size   int
	fulled chan struct{}
	l      sync.Mutex
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
	fulled := q.fulled
	if fulled != nil {
		<-fulled
	}

	q.l.Lock()
	defer q.l.Unlock()

	q.data.PushBack(b)

	if q.size > 0 && q.data.Len() == q.size {
		q.fulled = make(chan struct{})
	}
	return
}

// Dequeue 出队
func (q *Queue) Dequeue() (b *bytebuffer.ByteBuffer) {
	q.l.Lock()
	defer q.l.Unlock()

	front := q.data.Front()
	if front == nil {
		return nil
	}

	v := q.data.Remove(front)
	b = v.(*bytebuffer.ByteBuffer)

	if q.fulled != nil {
		close(q.fulled)
		q.fulled = nil
	}

	return
}
