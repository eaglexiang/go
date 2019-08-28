/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-06 17:20:27
 * @LastEditTime: 2019-08-28 20:20:35
 */

package smartstrmap

import "strings"

// SmartStrMap 对 map[string] interface{} 的封装
// string格式的key会被忽略大小写
type SmartStrMap struct {
	data map[string]interface{} // 存放的数据
}

// Set 设置 key - value 键值对
func (ssm *SmartStrMap) Set(key string, value interface{}) {
	if ssm.data == nil {
		ssm.data = make(map[string]interface{})
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
