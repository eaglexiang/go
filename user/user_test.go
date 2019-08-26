/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-06 18:13:57
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-03-03 04:56:13
 */

package user

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_UserStr(t *testing.T) {
	userStr := "testID:testPSWD"
	u, err := ParseValidUser(userStr)
	if err != nil {
		t.Error(err)
	}
	if str := u.ToString(); str != userStr {
		t.Error("User.ToString error: ", str)
	}

	userStr = "testID:testKey:100:3"
	u, err = ParseValidUser(userStr)
	if err != nil {
		t.Error(err)
	}
	if str := u.ToString(); str == userStr {
		t.Error("User.ToString error: ", str)
	}
	if u.ID != "testID" {
		t.Error("User.ID: ", u.ID)
	}
}

func Test_UserCheck(t *testing.T) {
	u, _ := ParseValidUser("test:testKey:100:3")
	reqUser, _ := ParseReqUser("testID:testKey", "127.0.0.1")
	err := u.CheckAuth(reqUser)
	if err != nil {
		t.Error(err)
	}
	reqUser.Location = "192.168.0.1"
	err = u.CheckAuth(reqUser)
	if err != nil {
		t.Error(err)
	}
	reqUser.Location = "192.168.0.2"
	err = u.CheckAuth(reqUser)
	if err != nil {
		t.Error(err)
	}
	reqUser.Location = "192.168.0.3"
	err = u.CheckAuth(reqUser)
	if err == nil {
		t.Error("should abort for 'too much loginstatus'")
	}
}

func Test_User_Concurrence(t *testing.T) {
	validUser, _ := ParseValidUser("testId:testKey::3")
	var countOfReqUsers = 100000
	var reqUsers []*ReqUser
	for i := 0; i < countOfReqUsers; i++ {
		location := strconv.FormatInt(int64(i), 10)
		reqUser, _ := ParseReqUser("testId:testKey", location)
		reqUsers = append(reqUsers, reqUser)
	}

	wg := sync.WaitGroup{}
	wg.Add(countOfReqUsers)
	for _, v := range reqUsers {
		// 测试是否触发 panic
		go validUser.checkAuthTest(v, func() { wg.Done() })
	}
	wg.Wait()
}

func Test_User_Count(t *testing.T) {
	validUser, _ := ParseValidUser("testId:testKey::5")
	for i := 0; i < 3; i++ {
		reqUser, _ := ParseReqUser("testId:testKey", fmt.Sprint(i))
		validUser.CheckAuth(reqUser)
	}
	count := validUser.Count()
	if count != FreeMsg {
		t.Error(count)
	}
}

func Test_User_SpeedLimit(t *testing.T) {
	user, _ := ParseValidUser("testID:testKey:100") // 100KB/s
	start := time.Now()
	for i := 0; i < 100; i++ {
		user.speedLimiter.Take() // 每秒100个tick
	}
	d := time.Since(start)
	if d < time.Millisecond*900 {
		t.Error("speed is too high")
	}
	if d > time.Millisecond*1100 {
		t.Error("speed is too low")
	}
}
