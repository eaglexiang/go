/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-17 15:52:23
 * @LastEditTime: 2019-08-26 23:39:18
 */

package logger

import (
	"fmt"
	"testing"

	"github.com/eaglexiang/go/settings"
)

func Test_logger(t *testing.T) {
	settings.Set("logger.debug", "off")
	print()
	settings.Set("logger.debug", "error")
	print()
	settings.Set("logger.debug", "warning")
	print()
	settings.Set("logger.debug", "info")
	print()
	settings.Set("logger.debug", "on")
	print()
}

func print() {
	fmt.Println("当前日志级别： " + settings.Get("logger.debug"))
	Error("测试错误")
	Warning("测试警告")
	Info("测试消息")
}

func Test_PrintFunc(t *testing.T) {
	Info(
		"测试函数输入（应该输入test）：",
		func() string {
			return "test"
		}(),
	)
}
