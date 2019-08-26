/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-06 17:30:28
 * @LastEditTime: 2019-08-26 23:31:44
 */

package settings

import (
	"errors"
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
	data   smartmap.SmartStrMap
}

// Set 配置
func (s *Settings) Set(key, value string) {
	className, ok := getClassName(key)
	if ok {
		child := s.GetChild(className)
		subKey, ok := getSubKey(key)
		if !ok {
			panic("no subKey for " + key) // 存在className时应该必然存在subKey
		}
		child.Set(subKey, value)
		return
	}

	s.data.Set(key, value)
}

// Get 获取配置，key不存在则触发panic
func (s Settings) Get(key string) (value string) {
	className, ok := getClassName(key)
	if ok {
		child := s.GetChild(className)
		subKey, ok := getSubKey(key)
		if !ok {
			panic("no subKey for " + key) // className与subKey必然同时存在
		}
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

// Bind 将bindKey与trueKey进行绑定，使bindKey单向映射为trueKey
// 一个trueKey可映射多个bindKey
func (s *Settings) Bind(bindKey string, trueKey string) {
	s.data.Bind(bindKey, trueKey)
}

// Exsit 判断 key 是否存在
func (s Settings) Exsit(key string) bool {
	_, ok := s.data.Get(key)
	return ok
}

// SetDefault 设置默认值，不会覆盖既有值
func (s *Settings) SetDefault(key, value string) {
	if s.Exsit(key) {
		return
	}
	s.Set(key, value)
}

// ImportLines 从 []string 导入数据
// 每行应该遵从以下格式：
// key = value
func (s Settings) ImportLines(data []string) error {
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

// ToString 输出为字符串
func (s Settings) ToString() string {
	var text string

	s.data.Range(func(k string, v interface{}) bool {
		text += k + " = " + v.(string) + "\n"
		return true
	})

	return text
}
