/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-07-24 21:22:36
 * @LastEditTime: 2019-07-24 21:22:36
 */

package tunnel

import (
	"sync"
)

type pipeStatus struct {
	l      sync.Mutex
	flowed bool // 管道已经开始流动
	closed bool // 管道已经关闭
}

func (s *pipeStatus) Clear() {
	s.flowed = false
	s.closed = false
}

func (s *pipeStatus) Close() {
	s.closed = true
}

func (s *pipeStatus) Closed() bool {
	return s.closed
}
