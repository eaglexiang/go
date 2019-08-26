/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-03 18:06:14
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-03-08 02:06:54
 */

package user

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const defaultLoginTTL = time.Minute * time.Duration(3) // 3min

const (
	// FullMsg 登录数量已满
	FullMsg = "full"
	// FreeMsg 登录数量未满
	FreeMsg = "free"
	// NoLimitMsg 没有登录限制
	NoLimitMsg = "no limit"
)

// loginStatus 用户的登录记录
// 登录记录的TTL为3分钟
type loginStatus struct {
	sync.Mutex
	sessions sync.Map // 所有登录实例
	count    int
	cap      int
}

// CreateloginStatus 创建loginStatus
// cap：同时登录地的限制数，cap>=0，默认为0，0表示不限制
func createloginStatus(cap ...int) *loginStatus {
	ls := loginStatus{}
	if len(cap) == 1 {
		if cap[0] < 0 {
			panic("Cap for loginStatus is less than 0")
		}
		ls.cap = cap[0]
	} else if len(cap) > 1 {
		panic("CreateloginStatus: too many args")
	}
	return &ls
}

// Login 登录
func (ls *loginStatus) login(id string) (err error) {
	if ls.cap == 0 {
		// 无限制状态，不需要登记
		return nil
	}

	var (
		s  interface{}
		ok bool
	)
	if s, ok = ls.sessions.Load(id); !ok {
		return ls.newSession(id)
	}
	s.(*session).fresh()
	return nil
}

func (ls *loginStatus) newSession(id string) (err error) {
	ls.Lock()
	defer ls.Unlock()
	err = ls.checkCount()
	if err != nil {
		ls.freshSessions()
		err = ls.checkCount()
	}
	if err != nil {
		return err
	}
	ls.sessions.Store(id, &session{
		updated: time.Now(),
		ttl:     defaultLoginTTL,
	})
	ls.count++
	return nil
}

func (ls *loginStatus) freshSessions() {
	ls.sessions.Range(func(key, value interface{}) bool {
		s := value.(*session)
		if !s.fresh() {
			ls.sessions.Delete(key)
			ls.count--
			return false
		}
		return true
	})
}

func (ls *loginStatus) checkCount() error {
	if ls.count == ls.cap {
		return errors.New("login full")
	} else if ls.count > ls.cap {
		panic("loginStatus.count > loginStatus.Cap")
	} else if ls.count < 0 {
		panic("loginStatus.count < 0")
	}
	return nil
}

func parseLoginCount(arg string) (int, error) {
	switch arg {
	case "private", "PRIVATE":
		return PrivateUser, nil
	case "share", "shared", "SHARED":
		return SharedUser, nil
	default:
		value, err := strconv.ParseInt(arg, 10, 32)
		return int(value), err
	}
}

// countStr 以比例的格式输出当前登录地占用状况，如 2/5
func (ls *loginStatus) countStr() string {
	if ls.cap == 0 {
		return NoLimitMsg
	}
	if ls.cap == ls.count {
		return FullMsg
	}
	return FreeMsg
}
