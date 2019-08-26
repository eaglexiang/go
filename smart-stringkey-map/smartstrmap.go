/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-06 17:20:27
 * @LastEditTime: 2019-08-26 21:38:23
 */

package smartstrmap

import "strings"

// SmartStrMap 对 map[string] interface{} 的封装
// string格式的key会被忽略大小写
// 可对键名进行绑定
type SmartStrMap struct {
	data map[string]interface{} // 存放的数据
	bind map[string]string      // 绑定的键名
}

// Set 设置 key - value 键值对
func (ssm *SmartStrMap) Set(key string, value interface{}) {
	if ssm.data == nil {
		ssm.data = make(map[string]interface{})
	}
	if trueKey, ok := ssm.getTrueKey(key); ok {
		key = trueKey
	}
	// key 统一使用小写进行存储和比较
	key = strings.ToLower(key)
	ssm.data[key] = value
}

// Get 获取 key 对应的 value，第二个返回值（bool）表示元素是否存在
func (ssm SmartStrMap) Get(key string) (interface{}, bool) {
	if ssm.data == nil {
		return nil, false
	}
	if trueKey, ok := ssm.getTrueKey(key); ok {
		key = trueKey
	}
	// key 统一使用小写进行存储和比较
	key = strings.ToLower(key)
	value, ok := ssm.data[key]
	return value, ok
}

// Range 遍历
func (ssm SmartStrMap) Range(f func(key string, value interface{}) bool) {
	if ssm.data == nil {
		return
	}
	for k, v := range ssm.data {
		if !f(k, v) {
			break
		}
	}
}

// Bind 绑定
// 建立bindKey到trueKey的映射
// 一个bindKey只可映一个trueKey
// 一个trueKey可接受多个bindKey的映射
func (ssm *SmartStrMap) Bind(bindKey, trueKey string) {
	if ssm.bind == nil {
		ssm.bind = make(map[string]string)
	}
	bindKey = strings.ToLower(bindKey)
	trueKey = strings.ToLower(trueKey)
	ssm.bind[bindKey] = trueKey
}

// getTrueKey 获取bindKey对应的trueKey
// 第二个返回值表示该bindKey是否存在
func (ssm SmartStrMap) getTrueKey(bindKey string) (trueKey string, ok bool) {
	if ssm.bind == nil {
		return "", false
	}
	bindKey = strings.ToLower(bindKey)
	trueKey, ok = ssm.bind[bindKey]
	return
}
