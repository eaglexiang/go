<!--
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-21 16:44:03
 * @LastEditTime: 2019-08-03 20:08:58
 -->
 
# go-tunnel

供Go语言使用的全双工TCP隧道。可以用它对接两个TCP连接，使数据在其间透明流动。

## 特性

* 支持流量整形
* 支持加密
* 自带对象池

## 基本使用

```golang
t := GetTunnel // NewTunnel会自动从对象池中获取元素
t.Flow()       // 数据开始流动，此函数阻塞
t.Close()      // 关闭，这会中断数据的流动
PutTunnel(t)   // 归还Tunne
```

## 流量整形

```golang
Tunnel.SetLimitor(l go.uber.org/ratelimit.Limiter) // 设置限速器
```

## 加密

设置加密解密器

```golang
t.SetLeftC(c mycipher.Cipher)
t.SetRightC(c mycipher.Cipher)
// ...
```
