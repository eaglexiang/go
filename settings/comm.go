/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-26 21:51:53
 * @LastEditTime: 2019-08-28 20:04:12
 */

package settings

import (
	"strings"
)

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
