<!--
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-21 16:10:15
 * @LastEditTime: 2019-03-03 19:06:41
 -->
 
# go-cache

[![codebeat badge](https://codebeat.co/badges/9f8bca7f-f757-4010-b894-02380200beab)](https://codebeat.co/projects/github-com-eaglexiang-go-cache-master)

供Go语言使用的缓存结构

特性：

* 以interface{}实现，支持任意数据结构
* 根据自定义TTL（生存时间）自动化地销毁过期缓存元素
* 高效复用请求

## 高效复用请求

当新元素被添加，会被自动置为阻塞态，所有对其调用都会被阻塞，直到Update函数被第一次调用为止。适用于只需单次处理，多次使用的场景。例如DNS缓存。