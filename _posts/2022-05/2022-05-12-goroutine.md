---
layout: post
title: "wait goroutine & stop goroutine"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Golang
  - Web Programming
---
---

等待协程 (防止子协程还没结束，主协程就结束了)：
- time.Sleep(2*time.Second)
- unbuffered channel
- sync.WaitGroup

[](https://www.cnblogs.com/sparkdev/p/10917536.html)
  

结束协程 (如果一个goroutine一直跑停不掉，就内存泄漏了)：
- chan + select
- context.WithCancel+select //多个goroutine共用context，同时停止(cancel()取消此context及其子context）
[](https://www.flysnow.org/2017/05/12/go-in-action-go-context.html)