---
layout: post
title: "Network Is To find A Way To Your Heart"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Network
---

>  **身体是应用层中的 message，内衣是传输层中的 TCP 头，外套是网络层中 IP 头，帽子和鞋子分别是网络接口层的帧头和帧尾**

## 术语
HTTP 的传输单位则是`消息或报文(message)`，TCP 层的传输单位是`段(segment)`，IP 层的传输单位是`包(packet)`，网络接口层的传输单位是`帧(frame)`。但这些名词并没有什么本质的区分，可以统称为`数据包`。

## 传输层
- 当传输层的数据包大小超过`MSS`(TCP 最大报文段长度)，就要将数据包分段

## 网络层
- 如果 IP 报文大小超过`MTU`(以太网中一般为 1500 字节)就会再次进行分包

- `IP`=`网络号`(子网掩码255)+`主机号`(子网掩码0)

- 路由算法找目标IP子网

## 网络接口层 
- 以太网（局域网）用MAC地址通信

- 如果目标IP不是本地局域网，则填入的MAC地址是路由器，也就是把数据包转发给路由器，路由器一直转发下一个路由器，直到转发到目标主机的路由器，发现 目标IP 是自己局域网内的主机，就会ARP请求获取目标主机的MAC地址

- 所以说，网络接口层主要为网络层提供「链路级别」传输的服务，负责在以太网、WiFi 这样的底层网络上发送原始数据包，工作在网卡这个层次。

- 数据包经路由器转发的过程中，源IP地址和目标IP地址是不会变的，源MAC地址和目标MAC地址是会变化的


## 一言以蔽之网络通信
`HTTP数据包` + `可靠传输的TCP头部` + `路由导航的IP头部` + `下一位置的MAC头部` + `路由器(交换机)`

![](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/%E9%94%AE%E5%85%A5%E7%BD%91%E5%9D%80%E8%BF%87%E7%A8%8B/21.jpg)
  

[yyds](https://xiaolincoding.com/network)