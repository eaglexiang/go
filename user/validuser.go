/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2018-10-08 10:51:05
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-04-03 20:24:34
 */

package user

import (
	"errors"
	"strconv"
	"strings"

	"go.uber.org/ratelimit"
)

// ValidUser 提供基本和轻量的账户系统
// 必须使用 ParseValidUser 函数进行构造
type ValidUser struct {
	user
	loginlog     *loginStatus
	speedLimiter ratelimit.Limiter
}

// ParseValidUser 通过格式化的字符串构造新的ValidUser
// 标准格式化字符串为 id:password:speedlimit:logincount
func ParseValidUser(userStr string) (validUser *ValidUser, err error) {
	items := formatValidUserStr(userStr)
	if items == nil {
		return nil, errors.New("invalid user text")
	}
	validUser = &ValidUser{
		user: user{
			ID:       items[0],
			Password: items[1],
		},
	}
	// 设置限速
	speedLimit, err := strconv.ParseInt(items[2], 10, 64)
	if err != nil {
		return nil, errors.New("when parse ValidUser: " + err.Error())
	}
	if speedLimit < 0 {
		return nil, errors.New("speed limit for ValidUser must not be less than 0")
	}
	if speedLimit > 0 {
		limiter := ratelimit.New(int(speedLimit))
		validUser.speedLimiter = limiter
	}
	// 设置最大同时登录地
	maxLoginCount := 0
	maxLoginCount, err = parseLoginCount(items[3])
	if err != nil {
		return nil, err
	}
	if maxLoginCount < 0 {
		return nil, errors.New("ParseValidUser: maxLoginCount < 0")
	}
	validUser.loginlog = createloginStatus(maxLoginCount)
	return
}

// 将ValidUser字符串进行分割、检查、格式化以及补全
// :key:speed:logins -> nil                无ID非法
// id::speed:logins  -> nil                无密码非法
// id:key:speed      -> [id,key,speed,0]
// id:key            -> [id,key,0,0]
// id:key::logins    -> [id,key,0,logins]
func formatValidUserStr(old string) []string {
	items := strings.Split(old, ":")
	if len(items) < 2 {
		return nil
	}
	if items[0] == "" || items[1] == "" {
		return nil
	}
	if len(items) == 2 {
		items = append(items, "0")
	}
	if items[2] == "" {
		items[2] = "0"
	}
	if len(items) == 3 {
		items = append(items, "0")
	}
	if items[3] == "" {
		items[3] = "0"
	}
	return items
}

// ToString 将ValidUser格式化为字符串
// 只输出ID和密码
func (user *ValidUser) ToString() string {
	return user.ID + ":" + user.Password
}

// CheckAuth 检查请求ValidUser的密码是否正确，并检查是否超出登录限制
func (user *ValidUser) CheckAuth(user2Check *ReqUser) error {
	if user.Password != user2Check.Password {
		return errors.New("ValidUser.CheckAuth -> incorrent username or password")
	}

	return user.loginlog.login(user2Check.Location)
}

// checkAuthTest 供测试使用的CheckAuth
// 会在CheckAuth退出后调用回调
func (user *ValidUser) checkAuthTest(user2Check *ReqUser, cb func()) {
	user.CheckAuth(user2Check)
	cb()
}

// SpeedLimiter 该用户的速度控制器
func (user *ValidUser) SpeedLimiter() ratelimit.Limiter {
	return user.speedLimiter
}

// Count 以比例的格式输出当前登录地占用状况，如 2/5
func (user *ValidUser) Count() string {
	return user.loginlog.countStr()
}
