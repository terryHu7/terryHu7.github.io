---
layout: post
title: "listen(), 3-way shake and then accept()"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Golang
  - Web Programming
---
---

在 TCP 三次握手过程中，当服务器收到客户端的 SYN 包后，内核会把该连接存储到半连接队列。

然后再向客户端发送 SYN+ACK 包，接着客户端会返回 ACK，服务端收到第三次握手的 ACK 后，内核会把连接从半连接队列移除，然后创建新的完全的连接，并将其增加到全连接队列 ，等待进程调用 `accept()` 把连接取出来。

全连接队列指的是服务器与客户端完了 TCP 三次握手后，还没有被 accept() 系统调用取走连接的队列。

![accept()](https://cdn.jsdelivr.net/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8D%8A%E8%BF%9E%E6%8E%A5%E5%92%8C%E5%85%A8%E8%BF%9E%E6%8E%A5/3.jpg)
