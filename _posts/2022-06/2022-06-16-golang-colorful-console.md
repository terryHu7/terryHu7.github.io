---
layout: post
title: "colorful fmt.Println in terminal"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Golang
---

```go
var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

fmt.Println("\033[1;31m**Red**\033[0m")
```  
As in regular expression, `\033`->`^`, `m`->`$`

As in markdown, `1`->`**bold**`
