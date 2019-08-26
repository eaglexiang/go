/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-03-17 15:22:44
 * @LastEditTime: 2019-06-14 22:23:50
 */

package logger

import (
	"log"
)

// 日志信息级别
const (
	NoLogType = iota
	ErrorLogType
	WarningLogType
	InfoLogType
	InvalidLogType // 非法的日志类型
)

var (
	logText  []string
	thegrade uint
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
	logText = append(logText, "off")
	logText = append(logText, "error")
	logText = append(logText, "warning")
	logText = append(logText, "info")
}

// printLog 打印日志
func printLog(grade uint, v ...interface{}) {
	if grade >= InvalidLogType {
		panic("log grade is invalid")
	}

	gradenow := gradeNow()
	if grade > gradenow {
		return
	}

	gradeName := logType2LogName(grade)

	out := []interface{}{gradeName + ": "}
	for _, _v := range v {
		switch t := _v.(type) {
		case func() string:
			out = append(out, t())
		default:
			out = append(out, t)
		}
	}
	log.Println(out)
}

// Error 错误日志
func Error(v ...interface{}) {
	printLog(ErrorLogType, v...)
}

// Warning 警告日志
func Warning(v ...interface{}) {
	printLog(WarningLogType, v...)
}

// Info 消息日志
func Info(v ...interface{}) {
	printLog(InfoLogType, v...)
}

func logType2LogName(grade uint) string {
	return logText[grade]
}

func logName2LogType(grade string) uint {
	var i uint
	for ; i < InvalidLogType; i++ {
		if logText[i] == grade {
			return i
		}
	}
	if grade == "on" {
		return InfoLogType
	}
	return InvalidLogType
}

// SetGrade 设置等级
func SetGrade(grade string) {
	thegrade = logName2LogType(grade)
}

func gradeNow() uint {
	return thegrade
}
