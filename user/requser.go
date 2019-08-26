/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-03 18:07:15
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-02-21 14:33:27
 */

package user

import (
	"errors"
	"strings"
)

// ReqUser 请求登录使用的临时用户
type ReqUser struct {
	user
	Location string
}

// ParseReqUser 通过字符串创建ReqUser
func ParseReqUser(userStr, locaction string) (*ReqUser, error) {
	items := strings.Split(userStr, ":")
	if len(items) < 2 {
		return nil, errors.New("invalid user text")
	}
	if items[0] == "" {
		return nil, errors.New("null username")
	}
	if items[1] == "" {
		return nil, errors.New("null password")
	}
	user := ReqUser{
		user: user{
			ID:       items[0],
			Password: items[1],
		},
		Location: locaction,
	}
	return &user, nil
}
