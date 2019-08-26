/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-26 21:53:56
 * @LastEditTime: 2019-08-26 23:23:30
 */

package settings

// Set 设置全局默认配置中 key 的值
func Set(key, value string) {
	globalSettings.Set(key, value)
}

// Get 获取全局默认配置中 key 的值
func Get(key string) string {
	return globalSettings.Get(key)
}

// GetInt64 获取全局默认配置中 key 的值，并转换为int64,转换失败会触发panic
func GetInt64(key string) int64 {
	return globalSettings.GetInt64(key)
}

// Exsit 判断全局默认配置中是否存在 key
func Exsit(key string) bool {
	return globalSettings.Exsit(key)
}

// SetDefault 设置全局默认配置中 key 的默认值
// 不会覆盖既有值
func SetDefault(key, value string) {
	globalSettings.SetDefault(key, value)
}

// ImportLines 从 []string 向全局默认配置中导入数据
func ImportLines(data []string) error {
	return globalSettings.ImportLines(data)
}

// ToString 将全局默认配置输出为格式化的字符串
func ToString() string {
	return globalSettings.ToString()
}

// Bind 将bindKey与trueKey进行绑定，使bindKey单向映射为trueKey
// 一个trueKey可映射多个bindKey
func Bind(bindKey, trueKey string) {
	globalSettings.Bind(bindKey, trueKey)
}

// GetChild 获取子类的Settings
func GetChild(className string) *Settings {
	return globalSettings.GetChild(className)
}
