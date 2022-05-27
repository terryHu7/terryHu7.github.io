---
layout: post
title: "wg.Done is wg.Add(-1)"
subtitle: 'Wait blocks until the WaitGroup counter is zero.'
author: "Terry"
header-style: text
tags:
  - Golang
---

```go
package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	capsbyte := "abcd.efg"
	strindex := strings.Index(capsbyte, "ef")
	fmt.Println(strindex)

	var wg sync.WaitGroup
	urls := []string{
		"www.google.com",
		"www.golang.org",
		"www.huty1998.github.io",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done() // wg.Add(-1)
			resp, err := http.Get(url)
			fmt.Println(resp, err)
			/*
				&{
					200 OK 200 HTTP/1.1 1 1
					map[Content-Type:[text/html;charset=utf-8]
					Date:[Wed, 23 Mar 2022 06:47:24 GMT]
					Expires:[Thu, 01 Jan 1970 00:00:00 GMT]
					Server:[Apache/2.4.29 (Ubuntu)]
					Set-Cookie:[XSRF_TOKEN=;Path=/;Expires=Thu, 01-Jan-1970 00:00:00 GMT;Max-Age=0]
					Vary:[Accept-Encoding]] 0xc000190040 -1 [] false true map[] 0xc00011a000 <nil>
				}
				<nil>
			*/
		}(url)
	}
	wg.Wait() // Wait blocks until the WaitGroup counter is zero.
}
```