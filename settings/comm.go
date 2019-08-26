/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-26 21:51:53
 * @LastEditTime: 2019-08-26 21:59:11
 */

package settings

import (
	"strings"
)

func getClassName(key string) (className string, ok bool) {
	items := strings.Split(key, classSep)
	if len(items) < 2 {
		return
	}

	className = items[0]
	ok = true
	return
}

func getSubKey(key string) (subKey string, ok bool) {
	items := strings.Split(key, classSep)
	if len(items) < 2 {
		return
	}

	subKey = strings.Join(items[1:], classSep)
	ok = true
	return
}
