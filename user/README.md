<!--
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-22 12:50:29
 * @LastEditTime: 2019-01-22 12:54:07
 -->
# go-user

供Go语言使用的轻量用户系统

单用户格式化字符串为`id:pswd:sl:lc`

四个字段的含义为：

字段名|含义
---|---
id|用户名
pswd|密码
sl|speed limit，该用户的限速（单位：KB/s）
lc|login count，该用户的同时登录地限制数