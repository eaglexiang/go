/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-05-17 21:38:12
 * @LastEditTime: 2019-05-17 22:20:01
 */

package counter

import (
	"sync/atomic"
)

// Counter 计数器
type Counter struct {
	Value int64
}

// Up 上升
func (c *Counter) Up(step ...int64) (result int64) {
	var _step int64 = 1
	if len(step) > 0 {
		_step = step[0]
	}

	result = atomic.AddInt64(&c.Value, _step)

	if c.Value < 0 {
		panic("integer overflow")
	}

	return
}

// Down 下降
func (c *Counter) Down(step ...int64) (result int64) {
	var _step int64 = 1
	if len(step) > 0 {
		_step = step[0]
	}

	result = atomic.AddInt64(&c.Value, -1*_step)

	if result < 0 {
		panic("counter value shouldn't be less than 0")
	}

	return
}
