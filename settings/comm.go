/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-26 21:51:53
 * @LastEditTime: 2019-09-08 11:41:46
 */

package settings

import (
	"strings"
)

// classSep 类分隔符
const classSep = "."

func getChild(key string) (className, subKey string, ok bool) {
	items := strings.Split(key, classSep)
	if len(items) < 2 {
		return
	}

	className = items[0]
	subKey = strings.Join(items[1:], classSep)
	ok = true
	return
}

// JoinClassNameAndKey 将class_name与sub_key用标准class分隔符连接起来
func JoinClassNameAndKey(className, subKey string) (key string) {
	key = className + classSep + subKey
	return
}
