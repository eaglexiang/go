/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-08 01:21:33
 * @LastEditTime: 2019-03-08 02:05:45
 */

package user

import "time"

type session struct {
	updated time.Time
	ttl     time.Duration
}

func (s *session) fresh() (alive bool) {
	now := time.Now()
	if now.Sub(s.updated) <= s.ttl {
		alive = true
	}
	s.updated = now
	return
}
