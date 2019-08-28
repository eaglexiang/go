/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-06 17:30:28
 * @LastEditTime: 2019-08-28 20:54:52
 */

package settings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	smartmap "github.com/eaglexiang/go/smart-stringkey-map"
)

// classSep 类分隔符
const classSep = "."

// globalSettings 默认的全局配置
var globalSettings Settings

// Settings 基于文本的配置管理器
// 忽略键名大小写
// 可设置键名绑定
type Settings struct {
	childs map[string]*Settings // map[class_name] class_settings
	binds  map[string]string    // map [bind_key] true_key
	data   smartmap.SmartStrMap
}

// Set 配置
func (s *Settings) Set(key, value string) {
	key = s.getTrueKey(key)

	className, subKey, ok := getChild(key)
	if ok {
		child := s.GetChild(className)
		child.Set(subKey, value)
		return
	}

	s.data.Set(key, value)
}

// Get 获取配置，key不存在则触发panic
func (s Settings) Get(key string) (value string) {
	key = s.getTrueKey(key)

	className, subKey, ok := getChild(key)
	if ok {
		child := s.GetChild(className)
		value = child.Get(subKey)
		return
	}

	_value, ok := s.data.Get(key)
	if !ok {
		panic("no key: " + key)
	}
	value = _value.(string)

	return
}

// GetInt64 获取配置，且尝试转换为int64,转换失败会触发panic
func (s Settings) GetInt64(key string) int64 {
	_v := s.Get(key)
	v, err := strconv.ParseInt(
		_v,
		10,
		32)

	if err != nil {
		panic(err)
	}

	return v
}

// Exsit 判断 key 是否存在
func (s Settings) Exsit(key string) bool {
	key = s.getTrueKey(key)

	className, subKey, ok := getChild(key)
	if ok {
		child := s.GetChild(className)
		return child.Exsit(subKey)
	}

	_, ok = s.data.Get(key)
	return ok
}

// SetDefault 设置默认值，不会覆盖既有值
func (s *Settings) SetDefault(key, value string) {
	key = s.getTrueKey(key)

	if s.Exsit(key) {
		return
	}
	s.Set(key, value)
}

// ImportLines 从 []string 导入数据
// 每行应该遵从以下格式：
// key = value
func (s *Settings) ImportLines(data []string) error {
	for _, line := range data {
		objs := strings.Split(line, "=")
		if len(objs) < 2 {
			return errors.New("invalid line: " + line)
		}
		key := strings.TrimSpace(objs[0])
		value := strings.Join(objs[1:], "=")
		value = strings.TrimSpace(value)
		s.Set(key, value)
	}
	return nil
}

// GetChild 获取子类的Settings
func (s *Settings) GetChild(className string) (child *Settings) {
	if s.childs == nil {
		s.childs = make(map[string]*Settings)
	}

	child, ok := s.childs[className]
	if !ok {
		child = new(Settings)
		s.childs[className] = child
	}
	return
}

func (s Settings) toSets() (sets []string) {
	s.data.Range(func(k string, v interface{}) bool {
		var set = fmt.Sprint(k, " = ", v)
		sets = append(sets, set)
		return true
	})

	return
}

// ToString 输出为字符串
func (s Settings) ToString() string {
	var sets = s.toSets()

	for className, child := range s.childs {
		var childSets = child.toSets()
		for i, v := range childSets {
			childSets[i] = fmt.Sprint(className, classSep, v)
		}
		sets = append(sets, childSets...)
	}

	var text = strings.Join(sets, "\n")

	return text
}

// Bind 绑定
// 建立bindKey到trueKey的映射
// 一个bindKey只可映一个trueKey
// 一个trueKey可接受多个bindKey的映射
func (s *Settings) Bind(bindKey, trueKey string) {
	if s.binds == nil {
		s.binds = make(map[string]string)
	}

	bindKey = strings.ToLower(bindKey)
	trueKey = strings.ToLower(trueKey)

	s.binds[bindKey] = trueKey
}

// getTrueKey 获取bindKey对应的trueKey，如果不存在Bind关系，则返回原bindKey
func (s Settings) getTrueKey(bindKey string) (trueKey string) {
	if s.binds == nil {
		trueKey = bindKey
		return
	}

	bindKey = strings.ToLower(bindKey)
	if key, ok := s.binds[bindKey]; ok {
		trueKey = key
	} else {
		trueKey = bindKey
	}
	return
}
